package gl_test

import (
	"testing"
)

// TODO(tmckee): reference the def, don't copy it
// Cribbed from /usr/include/GL/glew.h
const GLEW_OK = 0

func TestGlInit(t *testing.T) {
	//	require.Equal(t, gl.Init(), GLEW_OK)
}

/*
func TestCreateShader(t *testing.T) {
	require.Equal(t, GLEW_OK, gl.Init())

	shader := gl.CreateShader(gl.VERTEX_SHADER)
	if shader == 0 {
		t.Fail()
	}
	// require.NotEqual(t, 0, shader)
}
*/
