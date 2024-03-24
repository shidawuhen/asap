/*
*
@author: Jason Pang
@desc:
@date: 2024/3/24
*
*/
package sse

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Home() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.Query("user")
		topics := c.Query("topics")

		c.Writer.WriteString(fmt.Sprintf(`
			<!doctype html>
			<html lang="en">

			<head>
					<meta charset="UTF-8">
					<title>Server Sent Event</title>
			</head>

			<body>
			<div>Topic Messages:</div>
			<div class="event-data"></div>
			</body>

			<script src="https://code.jquery.com/jquery-1.11.1.js"></script>
			<script>
					// EventSource object of javascript listens the streaming events from our go server and prints the message.
					var stream = new EventSource("/sse/subscribe?user=%s&topics=%s");
					stream.addEventListener("message", function(e){
							$('.event-data').append(e.data + "</br>")
					});
                    stream.onerror = function(event) {
						stream.close();
					};
					
			</script>

			</html>`, user, topics))
	}
}
