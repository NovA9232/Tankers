#version 330

// Input vertex attributes (from vertex shader)
in vec2 fragTexCoord;
in vec4 fragColor;

// Input uniform values
uniform sampler2D texture0;
uniform vec4 colDiffuse;

// Output fragment color
out vec4 finalColor;

// NOTE: Add here your custom variables

uniform float time;

void main() {
	vec4 texelColor = texture(texture0, fragTexCoord);
	finalColor = vec4(abs(sin(time)) - texelColor.r, abs(cos(time)) - texelColor.g, texelColor.ba);
}
