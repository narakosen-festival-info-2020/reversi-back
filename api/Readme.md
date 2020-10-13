# API Documentation

## About Reversi

### /api/generate
Create the board.

**POST**
```json
{
    "board_type": "hoge"
}
```
normal type: "board"  
**Return**
```json
{
    "generate_time": "2020-10-13T02:52:41.791076185Z",
    "specific_code": "CzYWI8fMb2ookzqI"
}
```
Code "CzYWI8fMb2ookzqI" is Example

### /api/reversi/state
Get the current board information.  
**GET**  
Require http header: "Authorization: Bearer CzYWI8fMb2ookzqI"

**Return**
```json
{
    "board_type": "normal",
    "height": 8,
    "width": 8,
    "count_turn": 1,
    "who_turn": 1,
    "is_game_end": false,
    "can_place": true,
    "board":[
        [0, 0, 0, 0, 0, 0, 0, 0],
        [0, 0, 0, 0, 0, 0, 0, 0],
        [0, 0, 0, 0, 0, 0, 0, 0],
        [0, 0, 0, 1, 2, 0, 0, 0],
        [0, 0, 0, 2, 1, 0, 0, 0],
        [0, 0, 0, 0, 0, 0, 0, 0],
        [0, 0, 0, 0, 0, 0, 0, 0],
        [0, 0, 0, 0, 0, 0, 0, 0]
    ]
}
```

### /api/reversi/action
User action. (place stone)  
**POST**  
Require http header: "Authorization: Bearer CzYWI8fMb2ookzqI"  
```json
{
    "x": 0, 
    "y": 0
}
```

**Return**
```json
{
    "board_type": "normal",
    "height": 8,
    "width": 8,
    "count_turn": 1,
    "who_turn": 1,
    "is_game_end": false,
    "can_place": true,
    "board":[
        [0, 0, 0, 0, 0, 0, 0, 0],
        [0, 0, 0, 0, 0, 0, 0, 0],
        [0, 0, 0, 0, 0, 0, 0, 0],
        [0, 0, 0, 1, 2, 0, 0, 0],
        [0, 0, 1, 1, 1, 0, 0, 0],
        [0, 0, 0, 0, 0, 0, 0, 0],
        [0, 0, 0, 0, 0, 0, 0, 0],
        [0, 0, 0, 0, 0, 0, 0, 0]
    ]
}
```
