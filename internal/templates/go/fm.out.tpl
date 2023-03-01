{{ $inMessageName := .FieldMaskPairs.InMessage.Name }}
{{ $outMessageName := .FieldMaskPairs.OutMessage.Name }}
{{ $outMessagePkgName := .FieldMaskPairs.OutMessagePkgName }}
{{ $thisPackageName := .thisPkgName }}
{{ $fmFieldName := .FieldMaskPairs.FieldMaskField.Name.UpperCamelCase }}

{{ $messageName := .FieldMaskPairs.OutMessage.Name.UpperCamelCase }}

{{ if .FieldMaskPairs.GenOutMessageVar }}
// _fm_{{ $messageName }} is built in variable for {{ $messageName }} to call FieldMask.Append
{{ if eq $thisPackageName $outMessagePkgName }}
var _fm_Out_{{ $messageName }} = new({{ $messageName }})
{{ else }}
var _fm_Out_{{ $messageName }} = new({{if $outMessagePkgName }}{{ $outMessagePkgName }}.{{end}}{{ $messageName }})
{{ end }}
{{ end }}
{{ template "message" dict "Message" .FieldMaskPairs.OutMessage "inMessageName" $inMessageName "fmFieldName" $fmFieldName "inOut" "Out" "suffix" "" "pathSuffix" "" "messageName" $messageName }}

// Mask only affects the fields in the {{ $inMessageName }}.
{{ if eq $thisPackageName $outMessagePkgName }}
func (x *{{ $inMessageName }}_FieldMask) Mask(m *{{ $outMessageName }}) *{{ $outMessageName }} {
{{ else }}
func (x *{{ $inMessageName }}_FieldMask) Mask(m *{{if $outMessagePkgName }}{{ $outMessagePkgName }}.{{end}}{{ $outMessageName }}) *{{if $outMessagePkgName }}{{ $outMessagePkgName }}.{{end}}{{ $outMessageName }} {
{{ end }}
   switch x.maskMode {
   case pbfieldmask.MaskMode_Filter:
        x.mask.Filter(m)
   case pbfieldmask.MaskMode_Prune:
        x.mask.Prune(m)
   }

   return m
}
