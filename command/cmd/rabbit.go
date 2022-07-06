package cmd

import (
	"Walker/app/services"
	"Walker/app/utils/logger"
	"Walker/global"
	"github.com/spf13/cobra"
	"log"
)

var name string
var rabbitCmd = &cobra.Command{
	Use:   "rabbit",
	Short: "rabbit客户端",
	Long:  "根据指令启动,停止，重启rabbit客户端（hello）",
	Run: func(cmd *cobra.Command, args []string) {
		//启动一个rabbit客户端
		if len(args) != 1 {
			log.Printf("请传入关键词表明你需要启动的服务")
			return
		}
		switch args[0] {
		case "hello":
			RabbitStart()
		case "work_queue":
			//RabbitWorkQueueStart()
		default:
			log.Printf("你输入的命令无法判断想要做什么")

		}

	},
}

func init() {
	rootCmd.AddCommand(rabbitCmd)
}

func RabbitStart() {
	rabbit := services.RabbitService{}
	queue, channel := rabbit.Conn()
	msgs, err := channel.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		logger.InitLogger().Error(err.Error())
	}
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			global.Logger.Error("Received a message:" + string(d.Body))
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func RabbitWorkQueueStart() {

}
