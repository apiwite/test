package main

import (
	"log"
	"net/http"

	// "os" ===== aaaa
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/line/line-bot-sdk-go/v7/linebot/httphandler"
)

func main() {
	handler, err := httphandler.New(
		// os.Getenv("CHANNEL_SECRET"),
		// os.Getenv("CHANNEL_TOKEN"),
		"c6a6a734afca76bfccdc285ab7fb1a7a",
		"7QBsbHPpB10iMbRtr/6Rljy+cURi0cPp/Mluoc5BmDpLAlg+43N7mDT+s51V7jCOMyRvASBdk/u+9Lal8Xy0atRabGIizitkiCKJFBBOZ2QxtFNSKEb75WyBNdm9X7/qVy9cibexo6FAaJC2FF44WgdB04t89/1O/w1cDnyilFU=",
	)
	if err != nil {
		log.Fatal(err)
	}

	// Setup HTTP Server for receiving requests from LINE platform
	handler.HandleEvents(func(events []*linebot.Event, r *http.Request) {
		bot, err := handler.NewClient()
		if err != nil {
			log.Print(err)
			return
		}
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					if (message.Text == "Hello") || (message.Text == "hi") || (message.Text == "ดี") || (message.Text == "สวัสดี") {
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("hi :]")).Do(); err != nil {
							log.Print(err)
						}
					} else if message.Text == "pic" || (message.Text == "รูป") {
						url_umg := "https://us-fbcloud.net/wb/data/1314/1314125-img.uwdclx.45ktx.jpg"
						imageMessage := linebot.NewImageMessage(url_umg, url_umg)
						if _, err = bot.ReplyMessage(event.ReplyToken, imageMessage).Do(); err != nil {
							log.Print(err)
						}
					} else if message.Text == "ข้อมูลเพื่มเติม" || (message.Text == "info") {
						flexContainer, err := linebot.UnmarshalFlexMessageJSON([]byte(`{
							"type": "bubble",
							"size": "mega",
							"direction": "ltr",
							"header": {
							  "type": "box",
							  "layout": "horizontal",
							  "contents": [
								{
								  "type": "text",
								  "text": "🤖 BOTNOI",
								  "weight": "bold",
								  "size": "xl",
								  "color": "#00CAD3FF",
								  "contents": []
								}
							  ]
							},
							"hero": {
							  "type": "image",
							  "url": "https://www.img.in.th/images/43ac7365df713281d13c04cc2fe33f2b.png",
							  "size": "full",
							  "aspectRatio": "20:13",
							  "aspectMode": "cover",
							  "action": {
								"type": "uri",
								"label": "Action",
								"uri": "https://botnoigroup.com"
							  }
							},
							"body": {
							  "type": "box",
							  "layout": "horizontal",
							  "spacing": "md",
							  "contents": [
								{
								  "type": "box",
								  "layout": "vertical",
								  "flex": 1,
								  "contents": [
									{
									  "type": "image",
									  "url": "https://botnoigroup.com/wp-content/uploads/2020/08/botnoi_service-1.png",
									  "gravity": "bottom",
									  "size": "sm",
									  "aspectRatio": "4:3",
									  "aspectMode": "cover"
									},
									{
									  "type": "image",
									  "url": "https://botnoigroup.com/wp-content/uploads/2020/08/botnoi_service-2.png",
									  "margin": "md",
									  "size": "sm",
									  "aspectRatio": "4:3",
									  "aspectMode": "cover"
									}
								  ]
								},
								{
								  "type": "box",
								  "layout": "vertical",
								  "flex": 2,
								  "contents": [
									{
									  "type": "text",
									  "text": "Enterprise Chatbot Solution",
									  "size": "xs",
									  "flex": 1,
									  "gravity": "top",
									  "contents": []
									},
									{
									  "type": "text",
									  "text": "Read more...",
									  "size": "xs",
									  "color": "#0034B3FF",
									  "flex": 2,
									  "gravity": "top",
									  "action": {
										"type": "uri",
										"uri": "https://botnoigroup.com/chatbot/"
									  },
									  "contents": []
									},
									{
									  "type": "separator"
									},
									{
									  "type": "text",
									  "text": "Artificial Intelligence Solution",
									  "size": "xs",
									  "flex": 2,
									  "gravity": "center",
									  "contents": []
									},
									{
									  "type": "text",
									  "text": "Read more...",
									  "size": "xs",
									  "color": "#0034B3FF",
									  "flex": 1,
									  "gravity": "top",
									  "action": {
										"type": "uri",
										"uri": "https://botnoigroup.com/service/"
									  },
									  "contents": []
									}
								  ]
								}
							  ]
							},
							"footer": {
							  "type": "box",
							  "layout": "horizontal",
							  "contents": [
								{
								  "type": "button",
								  "action": {
									"type": "uri",
									"label": "Find More...🤖",
									"uri": "https://botnoigroup.com"
								  },
								  "color": "#004CFFFF",
								  "height": "sm"
								}
							  ]
							},
							"styles": {
							  "hero": {
								"backgroundColor": "#FFFFFFFF"
							  }
							}
						  }`))
						if err != nil {
							log.Print(err)
						}
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewFlexMessage("ข้อมูลเพื่มเติม", flexContainer)).Do(); err != nil {
							log.Print(err)
						}
					} else if message.Text == "ช่วยเหลือ" || (message.Text == "help") {
						flexContainer, err := linebot.UnmarshalFlexMessageJSON([]byte(`{
							"type": "bubble",
							"direction": "ltr",
							"hero": {
							  "type": "image",
							  "url": "https://botnoigroup.com/wp-content/uploads/2020/08/botnoi_service-1.png",
							  "size": "full",
							  "aspectRatio": "20:13",
							  "aspectMode": "cover",
							  "action": {
								"type": "uri",
								"label": "Action",
								"uri": "https://google.com"
							  }
							},
							"body": {
							  "type": "box",
							  "layout": "vertical",
							  "spacing": "md",
							  "action": {
								"type": "uri",
								"label": "Action",
								"uri": "https://linecorp.com"
							  },
							  "contents": [
								{
								  "type": "text",
								  "text": "Test-Bot command 🤖",
								  "weight": "bold",
								  "size": "xl",
								  "contents": []
								},
								{
								  "type": "box",
								  "layout": "vertical",
								  "spacing": "sm",
								  "contents": [
									{
									  "type": "box",
									  "layout": "baseline",
									  "contents": [
										{
										  "type": "text",
										  "text": "สวัสดี, ดี, hi, Hello",
										  "weight": "bold",
										  "size": "sm",
										  "margin": "sm",
										  "contents": []
										},
										{
										  "type": "text",
										  "text": "reply hi :]",
										  "size": "sm",
										  "color": "#AAAAAA",
										  "align": "end",
										  "contents": []
										}
									  ]
									},
									{
									  "type": "box",
									  "layout": "baseline",
									  "contents": [
										{
										  "type": "text",
										  "text": "รูป, pic",
										  "weight": "bold",
										  "size": "sm",
										  "flex": 0,
										  "margin": "sm",
										  "contents": []
										},
										{
										  "type": "text",
										  "text": "reply picture",
										  "size": "sm",
										  "color": "#AAAAAA",
										  "align": "end",
										  "contents": []
										}
									  ]
									},
									{
									  "type": "box",
									  "layout": "baseline",
									  "contents": [
										{
										  "type": "text",
										  "text": "ข้อมูลเพื่มเติม, info",
										  "weight": "bold",
										  "size": "sm",
										  "flex": 0,
										  "margin": "sm",
										  "contents": []
										},
										{
										  "type": "text",
										  "text": "reply infomation",
										  "size": "sm",
										  "color": "#AAAAAA",
										  "align": "end",
										  "contents": []
										}
									  ]
									}
								  ]
								},
								{
								  "type": "text",
								  "text": "---------.com",
								  "size": "xxs",
								  "color": "#AAAAAA",
								  "wrap": true,
								  "contents": []
								}
							  ]
							}
						  }`))
						if err != nil {
							log.Print(err)
						}
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewFlexMessage("ข้อมูลเพื่มเติม", flexContainer)).Do(); err != nil {
							log.Print(err)
						}
					} else {
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ขอโทษด้วย, ฉันไม่เข้าใจ ลองพิม help หรือ ช่วยเหลือ เพื่อดูคำสั่งที่รองรับ TT")).Do(); err != nil {
							log.Print(err)
						}
					}

				}
			}
		}
	})
	http.Handle("/callback", handler)
	// This is just a sample code.
	// For actually use, you must support HTTPS by using `ListenAndServeTLS`, reverse proxy or etc.
	// if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
	// 	log.Fatal(err)
	// }
	if err := http.ListenAndServe(":5678", nil); err != nil {
		log.Fatal(err)
	}
}

func Carousel_Template() {

}
