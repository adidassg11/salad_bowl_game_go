# salad_bowl_game_go
Game of Salad Bowl AKA Nouns AKA Celebrity

Rules - https://en.wikipedia.org/wiki/Celebrity_(game)

Milestones:
* Basic Gin + MySQL integration
* Gin API endpoints
* Gin endpoints to serve static content
* HTML + CSS (Bootstrap) to make a (very scrappy) frontend
* Gin endpoints that actually talk to frontend and DB. Can create games and words associated with them


Next steps
* Play with websockets to push events (20 min)
* Connect multiple clients to server and push events to all
* Websocket-based backend timer that disables parts of frontend


TODO list (bigger picture, not in order)
* expose to public web
* proper web hosting
* users, sessions, maybe keys/passwords (no strong auth for now)
* implement actual game logic
* fix the data models, requires ^^^
* try an ORM?
* move to different DB? Might depend on the web host
* learn debugging
* code cleanup / best practices


Notes to self
- The VSCode tool was somehow letting me break foreign key constraints, use different tool!
- Check out Render for hosting - https://gin-gonic.com/docs/deployment/ 