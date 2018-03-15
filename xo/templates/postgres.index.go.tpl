{{- $short := (shortname .Type.Name "err" "sqlstr" "db" "q" "res" "XOLog" .Fields) -}}
{{- $table := (schema .Schema .Type.Table.TableName) -}}

{{/* Prevent "[FUNCTION] redeclared in this block" errors
     See https://jira.nyu.edu/jira/browse/NYUP-397 */}}
{{- if (eq "hit_hit_name_7ea0901f_like" .Index.IndexName) -}}
{{- else -}}
    {{- if (eq "hit_hit_slug_b92dee21_like" .Index.IndexName) -}}
    {{- else -}}
        {{- if not (eq "lex_recognizer_replacer_0ef1a2a2_like" .Index.IndexName) }}

// {{ .FuncName }} retrieves a row from '{{ $table }}' as a {{ .Type.Name }}.
//
// Generated from index '{{ .Index.IndexName }}'.
func {{ .FuncName }}(db XODB{{ goparamlist .Fields true true }}) ({{ if not .Index.IsUnique }}[]{{ end }}*{{ .Type.Name }}, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`{{ colnames .Type.Fields }} ` +
		`FROM {{ $table }} ` +
		`WHERE {{ colnamesquery .Fields " AND " }}`

	// run query
	XOLog(sqlstr{{ goparamlist .Fields true false }})
{{- if .Index.IsUnique }}
	{{ $short }} := {{ .Type.Name }}{
	{{- if .Type.PrimaryKey }}
		_exists: true,
	{{ end -}}
	}

	err = db.QueryRow(sqlstr{{ goparamlist .Fields true false }}).Scan({{ fieldnames .Type.Fields (print "&" $short) }})
	if err != nil {
		return nil, err
	}

	return &{{ $short }}, nil
{{- else }}
	q, err := db.Query(sqlstr{{ goparamlist .Fields true false }})
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*{{ .Type.Name }}{}
	for q.Next() {
		{{ $short }} := {{ .Type.Name }}{
		{{- if .Type.PrimaryKey }}
			_exists: true,
		{{ end -}}
		}

		// scan
		err = q.Scan({{ fieldnames .Type.Fields (print "&" $short) }})
		if err != nil {
			return nil, err
		}

		res = append(res, &{{ $short }})
	}

	return res, nil
{{- end }}
}

        {{- end -}}
    {{- end -}}
{{- end -}}


