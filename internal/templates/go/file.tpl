// Code generated by protoc-gen-fieldmask. DO NOT EDIT.
// versions:
//  protoc-gen-fieldmask v0.4.1
// source: {{ .File.InputPath }}

package {{ pkg .File }}

import (
    pbfieldmask "github.com/yeqown/protoc-gen-fieldmask/proto/fieldmask"
    fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"

    {{ range $import := .ImportPaths }}
        {{ $import.PkgName}} "{{$import.ImportPath}}"
    {{ end }}
)

{{ range $idx, $pair := .FieldMaskPairs }}
    {{ $inGen := false }}
    {{ if $pair.FieldMaskOption.In }}
        {{ $inGen = $pair.FieldMaskOption.In.Gen }}
    {{ end }}
    {{ $outGen := false }}
    {{ if $pair.FieldMaskOption.Out }}
        {{ $outGen = $pair.FieldMaskOption.Out.Gen }}
    {{ end }}
    {{ if or (eq $inGen true) (eq $outGen true) }}
        {{ template "fm" . }}
    {{ end }}

    {{ if eq $inGen true }}
        {{ template "fm.in" . }}
    {{ end }}

    {{ if eq $outGen true }}
        {{ template "fm.out" . }}
    {{ end }}
{{ end }}
