package potofgreed

import (
	"fmt"
	"io"
	"sort"

	"github.com/go-yaml/yaml"
	"github.com/iancoleman/strcase"
	"golang.org/x/xerrors"
)

// Options is a configuration object used to generate the go files
type Options struct {
	Version       float32          `yaml:"version"`
	Package       string           `yaml:"package"`
	Models        map[string]Model `yaml:"models"`
	Relationships []Relationship   `yaml:"relationships"`
}

type Relationship struct {
	FromType  string `yaml:"from_type"`
	FromCount string `yaml:"from_count"`
	ToType    string `yaml:"to_type"`
	ToCount   string `yaml:"to_count"`
}

func Load(options io.Reader) (*Options, error) {
	o := &Options{
		Package: "models",
	}
	err := yaml.NewDecoder(options).Decode(o)
	if err != nil {
		return nil, xerrors.Errorf("couldn't decode options: %w", err)
	}
	return o, nil
}

func GenerateGoTypes(options *Options, out io.Writer) error {

	_, err := fmt.Fprintf(out, "type package %s\n\n", options.Package)
	if err != nil {
		return xerrors.Errorf("failed to write go package: %w", err)
	}
	type namedModel struct {
		Name  string
		Model Model
	}
	models := []namedModel{}
	for name, model := range options.Models {
		models = append(models, namedModel{name, model})
	}

	sort.Slice(models, func(i, j int) bool {
		return models[i].Name < models[j].Name
	})

	for _, rawModel := range models {
		goSrc, err := rawModel.Model.Golang()
		if err != nil {
			return xerrors.Errorf("failed to generate go definition for Raw%s: %w", rawModel.Name, err)
		}

		_, err = fmt.Fprintf(out, "type Raw%s %s\n", rawModel.Name, goSrc)
		if err != nil {
			return xerrors.Errorf("failed to write go source %s: %w", rawModel.Name, err)
		}

		model := Model{
			"": Type("Raw" + rawModel.Name).NotNull(),
		}

		for _, relation := range options.Relationships {
			if relation.FromType == rawModel.Name {
				model[relation.ToType] = Type(relation.ToType)
			}
			if relation.ToType == rawModel.Name {
				model[relation.FromType] = Type(relation.FromType)
			}
		}

		goSrc, err = model.Golang()
		if err != nil {
			return xerrors.Errorf("failed to generate go definition for %s: %w", rawModel.Name, err)
		}

		_, err = fmt.Fprintf(out, "type %s %s\n", rawModel.Name, goSrc)
		if err != nil {
			return xerrors.Errorf("failed to write go source %s: %w", rawModel.Name, err)
		}
	}
	return nil
}

func GenerateGraphQL(options *Options, out io.Writer) error {
	type namedModel struct {
		Name       string
		Model      Model
		InputModel Model
	}
	models := []namedModel{}
	for name, model := range options.Models {
		models = append(models, namedModel{name, model.Clone(), model})
	}

	sort.Slice(models, func(i, j int) bool {
		return models[i].Name < models[j].Name
	})

	fmt.Fprintf(out, `
schema {
	query: RootQuery
	mutation: RootMutation
}
type RootQuery {
`)

	for _, model := range models {
		fmt.Fprintf(out, "\t%s(id: ID!): %s\n", strcase.ToSnake(model.Name), model.Name)
		fmt.Fprintf(out, "\t%s_query(filter: %sFilter limit: Int! skip: Int): %s\n", strcase.ToSnake(model.Name), model.Name, model.Name)
	}
	fmt.Fprintf(out, `}
type RootMuttation {
`)
	for _, model := range models {
		fmt.Fprintf(out, "\tnew_%s(data: %sInput!): %s\n", strcase.ToSnake(model.Name), model.Name, model.Name)
		fmt.Fprintf(out, "\tupdate_%s(id: ID! data: %sInput!): %s\n", strcase.ToSnake(model.Name), model.Name, model.Name)
		fmt.Fprintf(out, "\tdelete_%s(id: ID!): %s\n", strcase.ToSnake(model.Name), model.Name)
	}
	fmt.Fprintf(out, "}\n")

	for _, model := range models {
		for _, relation := range options.Relationships {
			if relation.FromType == model.Name {
				model.Model[relation.ToType] = Type(relation.ToType)
			}
			if relation.ToType == model.Name {
				model.Model[relation.FromType] = Type(relation.FromType)
			}
		}

		for field, typ := range model.InputModel {
			model.InputModel[field] = typ.Nullable()
		}

		gqlSrc, err := model.InputModel.GraphQL()
		if err != nil {
			return xerrors.Errorf("failed to generate go definition for %s: %w", model.Name, err)
		}

		_, err = fmt.Fprintf(out, "input %sInput %s\n", model.Name, gqlSrc)
		if err != nil {
			return xerrors.Errorf("failed to write go source %s: %w", model.Name, err)
		}

		gqlSrc, err = model.Model.GraphQL()
		if err != nil {
			return xerrors.Errorf("failed to generate go definition for %s: %w", model.Name, err)
		}
		_, err = fmt.Fprintf(out, "type %s %s\n", model.Name, gqlSrc)
		if err != nil {
			return xerrors.Errorf("failed to write go source %s: %w", model.Name, err)
		}
	}
	return nil
}
