# Albion-KillBoard
Albion East Kill Death bot

## Dev Document
- Database:
    - Member ID, Name, Last time death, Last time kill
- Cronjob 1: Fetch members in guild every 1h
    - Save and add new member ID, Name
- Cronjob 2: Fetch kills / deaths for every player every 1m
    - API returns latest 10 kill/death
    - Compare 10 kills time with last time kill => Only process if kill is over this time
    - Put that KillDeathResponse to next process
    - Create `Goroutine` for each image process 