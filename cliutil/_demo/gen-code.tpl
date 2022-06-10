# run: kite gen parse cliutil/_demo/gen-code.tpl
colors = [red, blue, cyan, gray, green, yellow, magenta, info, warn, error ]

###

{{ foreach ($colors as $name) }}
    {{ $upName = ucfirst($name) }}
	cliutil.{{ $upName }}p("p:{{ $name }} color message, ")
	cliutil.{{ $upName }}f("f:%s color message, ", "{{ $name }}")
	cliutil.{{ $upName }}ln("ln:{{ $name }} color message print in cli.")
{{ endforeach }}