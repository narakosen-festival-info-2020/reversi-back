# reversi-back
Back-end of Reversi

## use
```sudo make build```  
```sudo make start```

## API
### /api/generate
#### POST
```
{
  "board_type": "normal"
}
```

#### return
```
response.data.specific_code > header
```

### /api/reversi/state
#### GET
```
{
  headers: {Authorization: `Bearer ${header}`,
}
```

#### return
```
response.data.board > board
```

### /api/revesi/state/action
#### POST
```
{
  "x": x,
  "y": y
}, {
  headers: {
    Authorization: `Bearer ${header}`,
  }
}
```

#### return
```
none
```
