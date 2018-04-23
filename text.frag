#version 300 es
out mediump vec4 FragColor;
in mediump vec2 TextCoord;

uniform sampler2D texture1;

void main()
{
    FragColor = texture(texture1, TextCoord);

}