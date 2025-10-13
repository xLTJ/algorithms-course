package main

import "testing"

func TestHashTable_Insert(t *testing.T) {
	type fields struct {
		table []*Entry
		size  int
		count int
	}
	type args struct {
		key   int
		value int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "insert into empty table",
			fields: fields{
				table: make([]*Entry, 11),
				size:  11,
				count: 0,
			},
			args:    args{key: 5, value: 100},
			wantErr: false,
		},
		{
			name: "insert with collision",
			fields: fields{
				table: func() []*Entry {
					table := make([]*Entry, 11)
					table[5] = &Entry{key: 5, value: 100, deleted: false}
					return table
				}(),
				size:  11,
				count: 1,
			},
			args:    args{key: 16, value: 200}, // 16 % 11 = 5, will collide
			wantErr: false,
		},
		{
			name: "update existing key",
			fields: fields{
				table: func() []*Entry {
					table := make([]*Entry, 11)
					table[5] = &Entry{key: 5, value: 100, deleted: false}
					return table
				}(),
				size:  11,
				count: 1,
			},
			args:    args{key: 5, value: 999},
			wantErr: false,
		},
		{
			name: "insert into deleted slot",
			fields: fields{
				table: func() []*Entry {
					table := make([]*Entry, 11)
					table[5] = &Entry{key: 5, value: 100, deleted: true}
					return table
				}(),
				size:  11,
				count: 0,
			},
			args:    args{key: 16, value: 200},
			wantErr: false,
		},
		{
			name: "insert into full table",
			fields: fields{
				table: func() []*Entry {
					table := make([]*Entry, 3)
					table[0] = &Entry{key: 0, value: 1, deleted: false}
					table[1] = &Entry{key: 1, value: 2, deleted: false}
					table[2] = &Entry{key: 2, value: 3, deleted: false}
					return table
				}(),
				size:  3,
				count: 3,
			},
			args:    args{key: 10, value: 100},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ht := &HashTable{
				table: tt.fields.table,
				size:  tt.fields.size,
				count: tt.fields.count,
			}
			if err := ht.Insert(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHashTable_Get(t *testing.T) {
	type fields struct {
		table []*Entry
		size  int
		count int
	}
	type args struct {
		key int
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantValue  int
		wantExists bool
	}{
		{
			name: "get existing key",
			fields: fields{
				table: func() []*Entry {
					table := make([]*Entry, 11)
					table[5] = &Entry{key: 5, value: 100, deleted: false}
					return table
				}(),
				size:  11,
				count: 1,
			},
			args:       args{key: 5},
			wantValue:  100,
			wantExists: true,
		},
		{
			name: "get non-existing key",
			fields: fields{
				table: make([]*Entry, 11),
				size:  11,
				count: 0,
			},
			args:       args{key: 5},
			wantValue:  0,
			wantExists: false,
		},
		{
			name: "get deleted key",
			fields: fields{
				table: func() []*Entry {
					table := make([]*Entry, 11)
					table[5] = &Entry{key: 5, value: 100, deleted: true}
					return table
				}(),
				size:  11,
				count: 0,
			},
			args:       args{key: 5},
			wantValue:  0,
			wantExists: false,
		},
		{
			name: "get key after collision",
			fields: fields{
				table: func() []*Entry {
					table := make([]*Entry, 11)
					table[5] = &Entry{key: 5, value: 100, deleted: false}
					table[6] = &Entry{key: 16, value: 200, deleted: false} // 16 % 11 = 5, probed to 6
					return table
				}(),
				size:  11,
				count: 2,
			},
			args:       args{key: 16},
			wantValue:  200,
			wantExists: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ht := &HashTable{
				table: tt.fields.table,
				size:  tt.fields.size,
				count: tt.fields.count,
			}
			gotValue, gotExists := ht.Get(tt.args.key)
			if gotValue != tt.wantValue {
				t.Errorf("Get() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
			if gotExists != tt.wantExists {
				t.Errorf("Get() gotExists = %v, want %v", gotExists, tt.wantExists)
			}
		})
	}
}

func TestHashTable_Delete(t *testing.T) {
	type fields struct {
		table []*Entry
		size  int
		count int
	}
	type args struct {
		key int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "delete existing key",
			fields: fields{
				table: func() []*Entry {
					table := make([]*Entry, 11)
					table[5] = &Entry{key: 5, value: 100, deleted: false}
					return table
				}(),
				size:  11,
				count: 1,
			},
			args:    args{key: 5},
			wantErr: false,
		},
		{
			name: "delete non-existing key",
			fields: fields{
				table: make([]*Entry, 11),
				size:  11,
				count: 0,
			},
			args:    args{key: 5},
			wantErr: true,
		},
		{
			name: "delete already deleted key (should not return error)",
			fields: fields{
				table: func() []*Entry {
					table := make([]*Entry, 11)
					table[5] = &Entry{key: 5, value: 100, deleted: true}
					return table
				}(),
				size:  11,
				count: 0,
			},
			args:    args{key: 5},
			wantErr: false,
		},
		{
			name: "delete key after collision",
			fields: fields{
				table: func() []*Entry {
					table := make([]*Entry, 11)
					table[5] = &Entry{key: 5, value: 100, deleted: false}
					table[6] = &Entry{key: 16, value: 200, deleted: false}
					return table
				}(),
				size:  11,
				count: 2,
			},
			args:    args{key: 16},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ht := &HashTable{
				table: tt.fields.table,
				size:  tt.fields.size,
				count: tt.fields.count,
			}
			if err := ht.Delete(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// Integration test to verify Insert -> Get -> Delete flow
func TestHashTable_Integration(t *testing.T) {
	ht := NewHashTable(11)

	// Insert some values
	if err := ht.Insert(5, 100); err != nil {
		t.Errorf("Insert failed: %v", err)
	}
	if err := ht.Insert(16, 200); err != nil {
		t.Errorf("Insert failed: %v", err)
	}

	// Verify they exist
	if val, exists := ht.Get(5); !exists || val != 100 {
		t.Errorf("Get(5) = %v, %v; want 100, true", val, exists)
	}
	if val, exists := ht.Get(16); !exists || val != 200 {
		t.Errorf("Get(16) = %v, %v; want 200, true", val, exists)
	}

	// Delete one
	if err := ht.Delete(5); err != nil {
		t.Errorf("Delete failed: %v", err)
	}

	// Verify it's gone
	if _, exists := ht.Get(5); exists {
		t.Errorf("Get(5) after delete should not exist")
	}

	// Verify the other still exists
	if val, exists := ht.Get(16); !exists || val != 200 {
		t.Errorf("Get(16) after deleting 5 = %v, %v; want 200, true", val, exists)
	}

	// Re-insert into deleted slot
	if err := ht.Insert(5, 999); err != nil {
		t.Errorf("Re-insert failed: %v", err)
	}
	if val, exists := ht.Get(5); !exists || val != 999 {
		t.Errorf("Get(5) after re-insert = %v, %v; want 999, true", val, exists)
	}
}
