package apirules

import (
	"bytes"
	"context"
	"fmt"
	"github.com/NubeIO/nrule/pprint"
	"github.com/NubeIO/rubix-assist/pkg/logger"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"time"
)

type PDFResponse struct {
	Result []pingResult
	Error  string
}

type PdfBody struct {
	Input          []byte `json:"input" binding:"required"`
	WriteToHomeDir bool   `json:"write_to_home_dir"`
}

func (inst *Client) PDF(pdfBody *PdfBody) *PingResponse {

	//pprint.PrintJSON(pdfBody)

	fmt.Println(pdfBody.WriteToHomeDir)

	_, err := inst.convert(inst.CTX, pdfBody.Input, pdfBody.WriteToHomeDir)
	fmt.Println(err)
	if err != nil {
		//return nil
	}

	r := &PingResponse{
		Result: nil,
		Error:  errorString(err),
	}
	pprint.PrintJSON(r)
	return r
}

type PDFApplication struct {
	PandocPath     string
	UserHome       string
	PandocDataDir  string
	CommandTimeout time.Duration
}

// copied from
// https://github.com/NubeIO/md-to-pdf

var permWrite = os.FileMode(0600)
var permDir = os.FileMode(0755)

func (inst *Client) convert(ctx context.Context, inputFile []byte, writeToHomeDir bool) ([]byte, error) {
	tmpdir := path.Join(os.TempDir(), fmt.Sprintf("pandocserver_%s", randStringRunes(10)))

	if err := os.Mkdir(tmpdir, permDir); err != nil {
		return nil, fmt.Errorf("could not create dir %q: %w", tmpdir, err)
	}
	defer os.RemoveAll(tmpdir)

	inputFileName := filepath.Join(tmpdir, fmt.Sprintf("%s.md", randStringRunes(10)))
	if err := os.WriteFile(inputFileName, inputFile, permWrite); err != nil {
		return nil, fmt.Errorf("could not create inputfile: %w", err)
	}

	outputDir := path.Join(tmpdir, "output")
	if err := os.Mkdir(outputDir, permDir); err != nil {
		return nil, fmt.Errorf("could not create output directory: %w", err)
	}
	pdfFilename := fmt.Sprintf("%s.pdf", randStringRunes(10))
	outputFilename := filepath.Join(outputDir, pdfFilename)

	args := []string{
		inputFileName,
		fmt.Sprintf("--output=%s", outputFilename),
		fmt.Sprintf("--data-dir=%s", inst.PdfApplication.PandocDataDir),
		"--from=markdown+yaml_metadata_block+raw_html+emoji",
	}
	commandCtx, cancel := context.WithTimeout(ctx, inst.PdfApplication.CommandTimeout)
	defer cancel()

	var out bytes.Buffer
	var stderr bytes.Buffer
	logger.Infof("path: %s", inst.PdfApplication.PandocPath)
	logger.Infof("going to call pandoc with the following args: %v", args)
	cmd := exec.CommandContext(commandCtx, "pandoc", args...)
	cmd.Dir = tmpdir
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		inst.killProcessIfRunning(cmd)
		return nil, fmt.Errorf("could not execute command %w: %s", err, stderr.String())
	}

	inst.killProcessIfRunning(cmd)

	content, err := os.ReadFile(outputFilename)
	if err != nil {
		return nil, fmt.Errorf("could not read output file: %w", err)
	}
	pdfFile := fmt.Sprintf("%s/%s", inst.PdfApplication.UserHome, pdfFilename)
	if writeToHomeDir {
		logger.Infof("write pdf to dir: %s", pdfFile)
		err := ioutil.WriteFile(pdfFile, content, 0644)
		if err != nil {
			fmt.Println(err, 4444444)
			logger.Errorf("write pdf to dir: %s err:%s", pdfFile, err.Error())
			return nil, err
		}
	}

	return content, nil
}

func (inst *Client) convertFromRead(ctx context.Context, inputFile string) (bool, error) {
	pdfFilename := fmt.Sprintf("%s.pdf", randStringRunes(10))
	outputFilename := filepath.Join(inst.PdfApplication.UserHome, pdfFilename)
	args := []string{ // works with MD and images
		inputFile,
		fmt.Sprintf("--output=%s", outputFilename),
		fmt.Sprintf("--data-dir=%s", inst.PdfApplication.PandocDataDir),
		"--from=markdown-implicit_figures",
	}

	commandCtx, cancel := context.WithTimeout(ctx, inst.PdfApplication.CommandTimeout)
	defer cancel()

	var out bytes.Buffer
	var stderr bytes.Buffer
	logger.Infof("path: %s", inst.PdfApplication.PandocPath)
	logger.Infof("going to call pandoc with the following args: %v", args)
	cmd := exec.CommandContext(commandCtx, "pandoc", args...)
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		inst.killProcessIfRunning(cmd)
		return false, fmt.Errorf("could not execute command %w: %s", err, stderr.String())
	}
	inst.killProcessIfRunning(cmd)
	return true, nil
}

func (inst *Client) killProcessIfRunning(cmd *exec.Cmd) {
	if cmd.Process == nil {
		return
	}
	if err := cmd.Process.Release(); err != nil {
		return
	}
	if err := cmd.Process.Kill(); err != nil {
		return
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
