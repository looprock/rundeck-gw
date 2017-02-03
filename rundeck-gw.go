package main

import (
    "strings"
    "time"
    "gopkg.in/gin-gonic/gin.v1"
    rundeck "github.com/lusis/go-rundeck/src/rundeck.v17"
)

func main() {
  router := gin.Default()

  router.GET("/hi", func(c *gin.Context) {
    c.JSON(200, gin.H{"status":  "hello world"})
  })

  router.GET("/beta/db/:name", func(c *gin.Context) {
    argString := make([]string, 2)
    argString[0] = "-name"
    argString[1] = c.Param("name")
    client := rundeck.NewClientFromEnv()
    jobopts := rundeck.RunOptions{
        Arguments: strings.Join(argString, " "),
    }
    res, err := client.RunJob("xx-xx-xx-xx-xx", jobopts)
    if err != nil {
        c.JSON(500, gin.H{"status":  err})
    } else {
        for {
            indata, err := client.GetExecution(res.Executions[0].ID)
            if err == nil {
              if indata.Status == "running" {
                time.Sleep(5)
                continue;
              } else {
                break;
              }
            } else {
              c.JSON(500, gin.H{"status":  err})
            }
        }
        outdata, err := client.GetExecution(res.Executions[0].ID)
        if err == nil {
          var httpstatus int
          if outdata.Status != "succeeded" {
            httpstatus = 500
          } else {
            httpstatus = 200
          }
          c.JSON(httpstatus, gin.H{"jobid": res.Executions[0].ID, "status":  outdata.Status})
        } else {
          c.JSON(500, gin.H{"status":  err})
        }
    }
  })

  router.GET("/beta/ingress/private/:name", func(c *gin.Context) {
    argString := make([]string, 4)
    argString[0] = "-name"
    argString[1] = c.Param("name")
    argString[2] = "-type"
    // hardcoding this because we don't want to create public ones via api yet
    argString[3] = "private"
    client := rundeck.NewClientFromEnv()
    jobopts := rundeck.RunOptions{
        Arguments: strings.Join(argString, " "),
    }
    res, err := client.RunJob("yy-yy-yy-yy-yy", jobopts)
    if err != nil {
        c.JSON(500, gin.H{"status":  err})
    } else {
        for {
            indata, err := client.GetExecution(res.Executions[0].ID)
            if err == nil {
              if indata.Status == "running" {
                time.Sleep(5)
                continue;
              } else {
                break;
              }
            } else {
              c.JSON(500, gin.H{"status":  err})
            }
        }
        outdata, err := client.GetExecution(res.Executions[0].ID)
        if err == nil {
          var httpstatus int
          if outdata.Status != "succeeded" {
            httpstatus = 500
          } else {
            httpstatus = 200
          }
          c.JSON(httpstatus, gin.H{"jobid": res.Executions[0].ID, "status":  outdata.Status})
        } else {
          c.JSON(500, gin.H{"status":  err})
        }
    }
  })

  router.Run(":8080")
}
