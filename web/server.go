package web

import (
	"ATS-CALCULATOR/matcher"
	"ATS-CALCULATOR/parser"
	"net/http"
	//"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()

	router.POST("/upload", handleUpload)

	router.Run(":8080")
}

func handleUpload(c *gin.Context) {
	resume, _ := c.FormFile("resume")
	job, _ := c.FormFile("job")

	// Parse resume and job description
	var resumeText, jobText string
	if resume.Filename[len(resume.Filename)-4:] == ".pdf" {
		resumeText, _ = parser.ParsePDF(resume.Filename)
	} else {
		resumeText, _ = parser.ParseDOCX(resume.Filename)
	}

	if job.Filename[len(job.Filename)-4:] == ".pdf" {
		jobText, _ = parser.ParsePDF(job.Filename)
	} else {
		jobText, _ = parser.ParseDOCX(job.Filename)
	}

	score := matcher.CalculateScore(resumeText, jobText)

	c.JSON(http.StatusOK, gin.H{
		"score": score,
	})
}
