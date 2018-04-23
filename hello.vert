#version 300 es
layout (location = 0) in vec3 aPos;
layout (location = 1) in vec2 aTextCoord;

out mediump vec2 TextCoord;
void main()
{

    gl_Position = vec4(aPos.x, aPos.y, aPos.z, 1.0);
    TextCoord = aTextCoord;
}
