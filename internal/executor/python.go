package executor

import (
	"bytes"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kurokobo/dify-unsandboxed-sandbox/internal/models"
)

type Python3Executor struct{}

func (p *Python3Executor) Run(c *gin.Context, req models.Request) (models.Response, error) {
	fileName := "/tmp/" + uuid.New().String() + ".py"
	fullCode := req.Preload + "\n" + req.Code
	if err := os.WriteFile(fileName, []byte(fullCode), 0644); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return models.Response{}, err
	}

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd := exec.Command("python3", fileName)
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf

	if err := cmd.Run(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return models.Response{}, err
	}
	stdout := stdoutBuf.String()
	stderr := stderrBuf.String()
	exitCode := cmd.ProcessState.ExitCode()

	return models.Response{
		Code:    exitCode,
		Message: "success",
		Data: models.RunData{
			Stdout: stdout,
			Error:  stderr,
		},
	}, nil
}

func (p *Python3Executor) Dependencies() models.Response {
	return models.Response{
		Code:    0,
		Message: "success",
		Data: models.DependenciesData{
			Dependencies: []models.Dependency{
				{Name: "base64", Version: ""},
				{Name: "binascii", Version: ""},
				{Name: "collections", Version: ""},
				{Name: "datetime", Version: ""},
				{Name: "functools", Version: ""},
				{Name: "hashlib", Version: ""},
				{Name: "hmac", Version: ""},
				{Name: "httpx", Version: ""},
				{Name: "itertools", Version: ""},
				{Name: "jinja2", Version: ""},
				{Name: "json", Version: ""},
				{Name: "math", Version: ""},
				{Name: "operator", Version: ""},
				{Name: "os", Version: ""},
				{Name: "random", Version: ""},
				{Name: "re", Version: ""},
				{Name: "requests", Version: ""},
				{Name: "string", Version: ""},
				{Name: "sys", Version: ""},
				{Name: "time", Version: ""},
				{Name: "traceback", Version: ""},
				{Name: "uuid", Version: ""},
			},
		},
	}
}
