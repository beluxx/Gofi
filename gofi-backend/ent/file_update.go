// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"gofi/ent/file"
	"gofi/ent/predicate"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// FileUpdate is the builder for updating File entities.
type FileUpdate struct {
	config
	name            *string
	isDirectory     *bool
	size            *int
	addsize         *int
	extension       *string
	mime            *string
	path            *string
	lastModified    *int64
	addlastModified *int64
	content         *string
	predicates      []predicate.File
}

// Where adds a new predicate for the builder.
func (fu *FileUpdate) Where(ps ...predicate.File) *FileUpdate {
	fu.predicates = append(fu.predicates, ps...)
	return fu
}

// SetName sets the name field.
func (fu *FileUpdate) SetName(s string) *FileUpdate {
	fu.name = &s
	return fu
}

// SetIsDirectory sets the isDirectory field.
func (fu *FileUpdate) SetIsDirectory(b bool) *FileUpdate {
	fu.isDirectory = &b
	return fu
}

// SetNillableIsDirectory sets the isDirectory field if the given value is not nil.
func (fu *FileUpdate) SetNillableIsDirectory(b *bool) *FileUpdate {
	if b != nil {
		fu.SetIsDirectory(*b)
	}
	return fu
}

// SetSize sets the size field.
func (fu *FileUpdate) SetSize(i int) *FileUpdate {
	fu.size = &i
	fu.addsize = nil
	return fu
}

// SetNillableSize sets the size field if the given value is not nil.
func (fu *FileUpdate) SetNillableSize(i *int) *FileUpdate {
	if i != nil {
		fu.SetSize(*i)
	}
	return fu
}

// AddSize adds i to size.
func (fu *FileUpdate) AddSize(i int) *FileUpdate {
	if fu.addsize == nil {
		fu.addsize = &i
	} else {
		*fu.addsize += i
	}
	return fu
}

// SetExtension sets the extension field.
func (fu *FileUpdate) SetExtension(s string) *FileUpdate {
	fu.extension = &s
	return fu
}

// SetMime sets the mime field.
func (fu *FileUpdate) SetMime(s string) *FileUpdate {
	fu.mime = &s
	return fu
}

// SetPath sets the path field.
func (fu *FileUpdate) SetPath(s string) *FileUpdate {
	fu.path = &s
	return fu
}

// SetLastModified sets the lastModified field.
func (fu *FileUpdate) SetLastModified(i int64) *FileUpdate {
	fu.lastModified = &i
	fu.addlastModified = nil
	return fu
}

// SetNillableLastModified sets the lastModified field if the given value is not nil.
func (fu *FileUpdate) SetNillableLastModified(i *int64) *FileUpdate {
	if i != nil {
		fu.SetLastModified(*i)
	}
	return fu
}

// AddLastModified adds i to lastModified.
func (fu *FileUpdate) AddLastModified(i int64) *FileUpdate {
	if fu.addlastModified == nil {
		fu.addlastModified = &i
	} else {
		*fu.addlastModified += i
	}
	return fu
}

// SetContent sets the content field.
func (fu *FileUpdate) SetContent(s string) *FileUpdate {
	fu.content = &s
	return fu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (fu *FileUpdate) Save(ctx context.Context) (int, error) {
	if fu.name != nil {
		if err := file.NameValidator(*fu.name); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"name\": %v", err)
		}
	}
	if fu.extension != nil {
		if err := file.ExtensionValidator(*fu.extension); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"extension\": %v", err)
		}
	}
	if fu.mime != nil {
		if err := file.MimeValidator(*fu.mime); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"mime\": %v", err)
		}
	}
	if fu.path != nil {
		if err := file.PathValidator(*fu.path); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"path\": %v", err)
		}
	}
	if fu.content != nil {
		if err := file.ContentValidator(*fu.content); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"content\": %v", err)
		}
	}
	return fu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FileUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FileUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FileUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fu *FileUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   file.Table,
			Columns: file.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: file.FieldID,
			},
		},
	}
	if ps := fu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value := fu.name; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: file.FieldName,
		})
	}
	if value := fu.isDirectory; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  *value,
			Column: file.FieldIsDirectory,
		})
	}
	if value := fu.size; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: file.FieldSize,
		})
	}
	if value := fu.addsize; value != nil {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: file.FieldSize,
		})
	}
	if value := fu.extension; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: file.FieldExtension,
		})
	}
	if value := fu.mime; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: file.FieldMime,
		})
	}
	if value := fu.path; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: file.FieldPath,
		})
	}
	if value := fu.lastModified; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  *value,
			Column: file.FieldLastModified,
		})
	}
	if value := fu.addlastModified; value != nil {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  *value,
			Column: file.FieldLastModified,
		})
	}
	if value := fu.content; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: file.FieldContent,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// FileUpdateOne is the builder for updating a single File entity.
type FileUpdateOne struct {
	config
	id              int
	name            *string
	isDirectory     *bool
	size            *int
	addsize         *int
	extension       *string
	mime            *string
	path            *string
	lastModified    *int64
	addlastModified *int64
	content         *string
}

// SetName sets the name field.
func (fuo *FileUpdateOne) SetName(s string) *FileUpdateOne {
	fuo.name = &s
	return fuo
}

// SetIsDirectory sets the isDirectory field.
func (fuo *FileUpdateOne) SetIsDirectory(b bool) *FileUpdateOne {
	fuo.isDirectory = &b
	return fuo
}

// SetNillableIsDirectory sets the isDirectory field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableIsDirectory(b *bool) *FileUpdateOne {
	if b != nil {
		fuo.SetIsDirectory(*b)
	}
	return fuo
}

// SetSize sets the size field.
func (fuo *FileUpdateOne) SetSize(i int) *FileUpdateOne {
	fuo.size = &i
	fuo.addsize = nil
	return fuo
}

// SetNillableSize sets the size field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableSize(i *int) *FileUpdateOne {
	if i != nil {
		fuo.SetSize(*i)
	}
	return fuo
}

// AddSize adds i to size.
func (fuo *FileUpdateOne) AddSize(i int) *FileUpdateOne {
	if fuo.addsize == nil {
		fuo.addsize = &i
	} else {
		*fuo.addsize += i
	}
	return fuo
}

// SetExtension sets the extension field.
func (fuo *FileUpdateOne) SetExtension(s string) *FileUpdateOne {
	fuo.extension = &s
	return fuo
}

// SetMime sets the mime field.
func (fuo *FileUpdateOne) SetMime(s string) *FileUpdateOne {
	fuo.mime = &s
	return fuo
}

// SetPath sets the path field.
func (fuo *FileUpdateOne) SetPath(s string) *FileUpdateOne {
	fuo.path = &s
	return fuo
}

// SetLastModified sets the lastModified field.
func (fuo *FileUpdateOne) SetLastModified(i int64) *FileUpdateOne {
	fuo.lastModified = &i
	fuo.addlastModified = nil
	return fuo
}

// SetNillableLastModified sets the lastModified field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableLastModified(i *int64) *FileUpdateOne {
	if i != nil {
		fuo.SetLastModified(*i)
	}
	return fuo
}

// AddLastModified adds i to lastModified.
func (fuo *FileUpdateOne) AddLastModified(i int64) *FileUpdateOne {
	if fuo.addlastModified == nil {
		fuo.addlastModified = &i
	} else {
		*fuo.addlastModified += i
	}
	return fuo
}

// SetContent sets the content field.
func (fuo *FileUpdateOne) SetContent(s string) *FileUpdateOne {
	fuo.content = &s
	return fuo
}

// Save executes the query and returns the updated entity.
func (fuo *FileUpdateOne) Save(ctx context.Context) (*File, error) {
	if fuo.name != nil {
		if err := file.NameValidator(*fuo.name); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"name\": %v", err)
		}
	}
	if fuo.extension != nil {
		if err := file.ExtensionValidator(*fuo.extension); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"extension\": %v", err)
		}
	}
	if fuo.mime != nil {
		if err := file.MimeValidator(*fuo.mime); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"mime\": %v", err)
		}
	}
	if fuo.path != nil {
		if err := file.PathValidator(*fuo.path); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"path\": %v", err)
		}
	}
	if fuo.content != nil {
		if err := file.ContentValidator(*fuo.content); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"content\": %v", err)
		}
	}
	return fuo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FileUpdateOne) SaveX(ctx context.Context) *File {
	f, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return f
}

// Exec executes the query on the entity.
func (fuo *FileUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FileUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fuo *FileUpdateOne) sqlSave(ctx context.Context) (f *File, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   file.Table,
			Columns: file.Columns,
			ID: &sqlgraph.FieldSpec{
				Value:  fuo.id,
				Type:   field.TypeInt,
				Column: file.FieldID,
			},
		},
	}
	if value := fuo.name; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: file.FieldName,
		})
	}
	if value := fuo.isDirectory; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  *value,
			Column: file.FieldIsDirectory,
		})
	}
	if value := fuo.size; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: file.FieldSize,
		})
	}
	if value := fuo.addsize; value != nil {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: file.FieldSize,
		})
	}
	if value := fuo.extension; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: file.FieldExtension,
		})
	}
	if value := fuo.mime; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: file.FieldMime,
		})
	}
	if value := fuo.path; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: file.FieldPath,
		})
	}
	if value := fuo.lastModified; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  *value,
			Column: file.FieldLastModified,
		})
	}
	if value := fuo.addlastModified; value != nil {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  *value,
			Column: file.FieldLastModified,
		})
	}
	if value := fuo.content; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: file.FieldContent,
		})
	}
	f = &File{config: fuo.config}
	_spec.Assign = f.assignValues
	_spec.ScanValues = f.scanValues()
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return f, nil
}
