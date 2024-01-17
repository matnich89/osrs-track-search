# Osrs-Track Search

## Description
Osrs-Track Search is a dedicated service for fetching player statistics from the Old School RuneScape highscores. It's a part of the broader osrs-track service architecture. This service interfaces with the official RuneScape highscores [here](https://secure.runescape.com/m=hiscore_oldschool/overall).

Designed for flexibility, it can be deployed on Kubernetes or run locally for development and debugging purposes.

## APIs
The service currently supports queries for various RuneScape highscore types including Standard, Ironman, HCIM, and UIM. (Note: Currently, only the Ironman type is implemented.)

### **GET /ironman**
Fetches statistics for Ironman characters.

**Parameters:**
- `character`: The name of the Ironman character.

**Functionality:**
- Retrieves character scores from Jagex (the owner of RuneScape).
- Processes and returns the stats in JSON format.

**Response Codes:**
- `200 OK`: Successful retrieval of stats.
- `400 Bad Request`: Invalid or missing character name.
- `404 Not Found`: Character does not exist.
- `502 Bad Gateway`: Issues with the external service.
- `500 Internal Server Error`: Unexpected processing error.

**Example:**
- **Request**: `GET http://{host}:8080/ironman?character=P I C K LE`
- **Response**:
```json
{
  "character": "P I C K LE",
  "stats": [
    {
      "skill": "Overall",
      "rank": 104686,
      "level": 1763,
      "xp": 45742159
    },
    // Additional skill stats...
  ]
}
