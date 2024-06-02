// Code generated by file2byteslice. DO NOT EDIT.

// +build example

package main

var lighting_go = []byte("// Copyright 2020 The Ebiten Authors\n//\n// Licensed under the Apache License, Version 2.0 (the \"License\");\n// you may not use this file except in compliance with the License.\n// You may obtain a copy of the License at\n//\n//     http://www.apache.org/licenses/LICENSE-2.0\n//\n// Unless required by applicable law or agreed to in writing, software\n// distributed under the License is distributed on an \"AS IS\" BASIS,\n// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.\n// See the License for the specific language governing permissions and\n// limitations under the License.\n\n// +build ignore\n\npackage main\n\nvar Time float\nvar Cursor vec2\nvar ScreenSize vec2\n\nfunc Fragment(position vec4, texCoord vec2, color vec4) vec4 {\n\tlightpos := vec3(Cursor, 50)\n\tlightdir := normalize(lightpos - position.xyz)\n\tnormal := normalize(imageSrc1UnsafeAt(texCoord) - 0.5)\n\tambient := 0.25\n\tdiffuse := 0.75 * max(0.0, dot(normal.xyz, lightdir))\n\treturn imageSrc0UnsafeAt(texCoord) * (ambient + diffuse)\n}\n")