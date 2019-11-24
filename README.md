
# Minesweeper

My version!

## Indices

* [Default](#default)

  * [http://localhost:8080/minesweeper/status?game_id=ec30fb3d-5f7e-4bc4-9724-f39f0498e45c&token=07d47f41-ed9d-4ee5-95cc-c2699cab4abc](#1-https:minesweeperdoghelcomarminesweeperstatus?game_id=ec30fb3d-5f7e-4bc4-9724-f39f0498e45c&token=07d47f41-ed9d-4ee5-95cc-c2699cab4abc)
  * [http://localhost:8080/minesweeper/mark](#2-https:minesweeperdoghelcomarminesweepermark)
  * [http://localhost:8080/minesweeper/play](#3-https:minesweeperdoghelcomarminesweeperplay)
  * [http://localhost:8080/minesweeper/init](#4-https:minesweeperdoghelcomarminesweeperinit)
  * [http://localhost:8080/ping](#5-https:minesweeperdoghelcomarping)


--------


## Default



### 1. http://localhost:8080/minesweeper/status?game_id=ec30fb3d-5f7e-4bc4-9724-f39f0498e45c&token=07d47f41-ed9d-4ee5-95cc-c2699cab4abc



***Endpoint:***

```bash
Method: GET
Type: 
URL: http://localhost:8080/minesweeper/status
```



***Query params:***

| Key | Value | Description |
| --- | ------|-------------|
| game_id | ec30fb3d-5f7e-4bc4-9724-f39f0498e45c |  |
| token | 07d47f41-ed9d-4ee5-95cc-c2699cab4abc |  |



***Responses:***


Status: http://localhost:8080/minesweeper/status?game_id=ec30fb3d-5f7e-4bc4-9724-f39f0498e45c&token=07d47f41-ed9d-4ee5-95cc-c2699cab4abc | Code: 200



```js
{
    "game_id": "ec30fb3d-5f7e-4bc4-9724-f39f0498e45c",
    "token": "07d47f41-ed9d-4ee5-95cc-c2699cab4abc",
    "human_friendly_board": "[0][0][ ][ ][ ][ ][ ][ ][ ][ ]\n[0][0][ ][ ][ ][ ][ ][ ][ ][ ]\n[ ][ ][ ][ ][ ][ ][ ][ ][ ][ ]\n[ ][ ][ ][ ][ ][ ][ ][ ][ ][ ]\n[ ][ ][ ][ ][ ][ ][ ][ ][ ][ ]\n[ ][ ][ ][ ][ ][ ][ ][ ][ ][ ]\n[ ][ ][ ][ ][ ][ ][ ][ ][ ][ ]\n[ ][ ][ ][ ][ ][ ][ ][ ][ ][ ]\n[ ][ ][ ][ ][ ][ ][ ][ ][ ][ ]\n[ ][ ][ ][ ][ ][ ][ ][ ][ ][?]\n",
    "board": [
        [
            {
                "clicked": true,
                "label": "[0]",
                "X": 0,
                "Y": 0
            },
            {
                "clicked": false,
                "label": "[0]",
                "X": 1,
                "Y": 0
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 2,
                "Y": 0
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 3,
                "Y": 0
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 4,
                "Y": 0
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 5,
                "Y": 0
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 6,
                "Y": 0
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 7,
                "Y": 0
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 8,
                "Y": 0
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 9,
                "Y": 0
            }
        ],
        [
            {
                "clicked": false,
                "label": "[0]",
                "X": 0,
                "Y": 1
            },
            {
                "clicked": false,
                "label": "[0]",
                "X": 1,
                "Y": 1
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 2,
                "Y": 1
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 3,
                "Y": 1
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 4,
                "Y": 1
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 5,
                "Y": 1
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 6,
                "Y": 1
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 7,
                "Y": 1
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 8,
                "Y": 1
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 9,
                "Y": 1
            }
        ],
        [
            {
                "clicked": false,
                "label": "[ ]",
                "X": 0,
                "Y": 2
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 1,
                "Y": 2
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 2,
                "Y": 2
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 3,
                "Y": 2
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 4,
                "Y": 2
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 5,
                "Y": 2
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 6,
                "Y": 2
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 7,
                "Y": 2
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 8,
                "Y": 2
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 9,
                "Y": 2
            }
        ],
        [
            {
                "clicked": false,
                "label": "[ ]",
                "X": 0,
                "Y": 3
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 1,
                "Y": 3
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 2,
                "Y": 3
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 3,
                "Y": 3
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 4,
                "Y": 3
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 5,
                "Y": 3
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 6,
                "Y": 3
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 7,
                "Y": 3
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 8,
                "Y": 3
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 9,
                "Y": 3
            }
        ],
        [
            {
                "clicked": false,
                "label": "[ ]",
                "X": 0,
                "Y": 4
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 1,
                "Y": 4
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 2,
                "Y": 4
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 3,
                "Y": 4
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 4,
                "Y": 4
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 5,
                "Y": 4
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 6,
                "Y": 4
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 7,
                "Y": 4
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 8,
                "Y": 4
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 9,
                "Y": 4
            }
        ],
        [
            {
                "clicked": false,
                "label": "[ ]",
                "X": 0,
                "Y": 5
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 1,
                "Y": 5
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 2,
                "Y": 5
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 3,
                "Y": 5
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 4,
                "Y": 5
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 5,
                "Y": 5
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 6,
                "Y": 5
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 7,
                "Y": 5
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 8,
                "Y": 5
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 9,
                "Y": 5
            }
        ],
        [
            {
                "clicked": false,
                "label": "[ ]",
                "X": 0,
                "Y": 6
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 1,
                "Y": 6
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 2,
                "Y": 6
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 3,
                "Y": 6
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 4,
                "Y": 6
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 5,
                "Y": 6
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 6,
                "Y": 6
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 7,
                "Y": 6
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 8,
                "Y": 6
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 9,
                "Y": 6
            }
        ],
        [
            {
                "clicked": false,
                "label": "[ ]",
                "X": 0,
                "Y": 7
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 1,
                "Y": 7
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 2,
                "Y": 7
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 3,
                "Y": 7
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 4,
                "Y": 7
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 5,
                "Y": 7
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 6,
                "Y": 7
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 7,
                "Y": 7
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 8,
                "Y": 7
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 9,
                "Y": 7
            }
        ],
        [
            {
                "clicked": false,
                "label": "[ ]",
                "X": 0,
                "Y": 8
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 1,
                "Y": 8
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 2,
                "Y": 8
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 3,
                "Y": 8
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 4,
                "Y": 8
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 5,
                "Y": 8
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 6,
                "Y": 8
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 7,
                "Y": 8
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 8,
                "Y": 8
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 9,
                "Y": 8
            }
        ],
        [
            {
                "clicked": false,
                "label": "[ ]",
                "X": 0,
                "Y": 9
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 1,
                "Y": 9
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 2,
                "Y": 9
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 3,
                "Y": 9
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 4,
                "Y": 9
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 5,
                "Y": 9
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 6,
                "Y": 9
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 7,
                "Y": 9
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 8,
                "Y": 9
            },
            {
                "clicked": false,
                "label": "[?]",
                "X": 9,
                "Y": 9
            }
        ]
    ],
    "game_setup": {
        "cols": 10,
        "rows": 10,
        "mines": 10
    },
    "game_status": "ACTIVE",
    "start_timestamp": 1573968107,
    "elapsed_time": 125588
}
```



### 2. http://localhost:8080/minesweeper/mark



***Endpoint:***

```bash
Method: POST
Type: RAW
URL: http://localhost:8080/minesweeper/mark
```


***Headers:***

| Key | Value | Description |
| --- | ------|-------------|
| Content-Type | application/json |  |



***Body:***

```js        
{
    "game_id": "ec30fb3d-5f7e-4bc4-9724-f39f0498e45c",
    "token": "07d47f41-ed9d-4ee5-95cc-c2699cab4abc",
    "X": 9,
    "Y": 9,
    "mark": "?"
}
```



***Responses:***


Status: http://localhost:8080/minesweeper/mark | Code: 200



```js
{
    "game_id": "ec30fb3d-5f7e-4bc4-9724-f39f0498e45c",
    "token": "07d47f41-ed9d-4ee5-95cc-c2699cab4abc",
    "human_friendly_board": "[0][0][ ][ ][ ][ ][ ][ ][ ][ ]\n[0][0][ ][ ][ ][ ][ ][ ][ ][ ]\n[ ][ ][ ][ ][ ][ ][ ][ ][ ][ ]\n[ ][ ][ ][ ][ ][ ][ ][ ][ ][ ]\n[ ][ ][ ][ ][ ][ ][ ][ ][ ][ ]\n[ ][ ][ ][ ][ ][ ][ ][ ][ ][ ]\n[ ][ ][ ][ ][ ][ ][ ][ ][ ][ ]\n[ ][ ][ ][ ][ ][ ][ ][ ][ ][ ]\n[ ][ ][ ][ ][ ][ ][ ][ ][ ][ ]\n[ ][ ][ ][ ][ ][ ][ ][ ][ ][?]\n",
    "board": [
        [
            {
                "clicked": true,
                "label": "[0]",
                "X": 0,
                "Y": 0
            },
            {
                "clicked": false,
                "label": "[0]",
                "X": 1,
                "Y": 0
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 2,
                "Y": 0
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 3,
                "Y": 0
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 4,
                "Y": 0
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 5,
                "Y": 0
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 6,
                "Y": 0
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 7,
                "Y": 0
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 8,
                "Y": 0
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 9,
                "Y": 0
            }
        ],
        [
            {
                "clicked": false,
                "label": "[0]",
                "X": 0,
                "Y": 1
            },
            {
                "clicked": false,
                "label": "[0]",
                "X": 1,
                "Y": 1
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 2,
                "Y": 1
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 3,
                "Y": 1
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 4,
                "Y": 1
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 5,
                "Y": 1
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 6,
                "Y": 1
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 7,
                "Y": 1
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 8,
                "Y": 1
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 9,
                "Y": 1
            }
        ],
        [
            {
                "clicked": false,
                "label": "[ ]",
                "X": 0,
                "Y": 2
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 1,
                "Y": 2
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 2,
                "Y": 2
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 3,
                "Y": 2
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 4,
                "Y": 2
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 5,
                "Y": 2
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 6,
                "Y": 2
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 7,
                "Y": 2
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 8,
                "Y": 2
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 9,
                "Y": 2
            }
        ],
        [
            {
                "clicked": false,
                "label": "[ ]",
                "X": 0,
                "Y": 3
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 1,
                "Y": 3
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 2,
                "Y": 3
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 3,
                "Y": 3
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 4,
                "Y": 3
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 5,
                "Y": 3
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 6,
                "Y": 3
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 7,
                "Y": 3
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 8,
                "Y": 3
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 9,
                "Y": 3
            }
        ],
        [
            {
                "clicked": false,
                "label": "[ ]",
                "X": 0,
                "Y": 4
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 1,
                "Y": 4
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 2,
                "Y": 4
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 3,
                "Y": 4
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 4,
                "Y": 4
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 5,
                "Y": 4
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 6,
                "Y": 4
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 7,
                "Y": 4
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 8,
                "Y": 4
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 9,
                "Y": 4
            }
        ],
        [
            {
                "clicked": false,
                "label": "[ ]",
                "X": 0,
                "Y": 5
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 1,
                "Y": 5
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 2,
                "Y": 5
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 3,
                "Y": 5
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 4,
                "Y": 5
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 5,
                "Y": 5
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 6,
                "Y": 5
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 7,
                "Y": 5
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 8,
                "Y": 5
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 9,
                "Y": 5
            }
        ],
        [
            {
                "clicked": false,
                "label": "[ ]",
                "X": 0,
                "Y": 6
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 1,
                "Y": 6
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 2,
                "Y": 6
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 3,
                "Y": 6
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 4,
                "Y": 6
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 5,
                "Y": 6
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 6,
                "Y": 6
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 7,
                "Y": 6
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 8,
                "Y": 6
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 9,
                "Y": 6
            }
        ],
        [
            {
                "clicked": false,
                "label": "[ ]",
                "X": 0,
                "Y": 7
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 1,
                "Y": 7
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 2,
                "Y": 7
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 3,
                "Y": 7
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 4,
                "Y": 7
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 5,
                "Y": 7
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 6,
                "Y": 7
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 7,
                "Y": 7
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 8,
                "Y": 7
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 9,
                "Y": 7
            }
        ],
        [
            {
                "clicked": false,
                "label": "[ ]",
                "X": 0,
                "Y": 8
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 1,
                "Y": 8
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 2,
                "Y": 8
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 3,
                "Y": 8
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 4,
                "Y": 8
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 5,
                "Y": 8
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 6,
                "Y": 8
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 7,
                "Y": 8
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 8,
                "Y": 8
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 9,
                "Y": 8
            }
        ],
        [
            {
                "clicked": false,
                "label": "[ ]",
                "X": 0,
                "Y": 9
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 1,
                "Y": 9
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 2,
                "Y": 9
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 3,
                "Y": 9
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 4,
                "Y": 9
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 5,
                "Y": 9
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 6,
                "Y": 9
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 7,
                "Y": 9
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 8,
                "Y": 9
            },
            {
                "clicked": false,
                "label": "[?]",
                "X": 9,
                "Y": 9
            }
        ]
    ],
    "game_setup": {
        "cols": 10,
        "rows": 10,
        "mines": 10
    },
    "game_status": "ACTIVE",
    "start_timestamp": 1573968107,
    "elapsed_time": 125588
}
```



### 3. http://localhost:8080/minesweeper/play



***Endpoint:***

```bash
Method: POST
Type: RAW
URL: http://localhost:8080/minesweeper/play
```


***Headers:***

| Key | Value | Description |
| --- | ------|-------------|
| Content-Type | application/json |  |



***Body:***

```js        
{
    "game_id": "ec30fb3d-5f7e-4bc4-9724-f39f0498e45c",
    "token": "07d47f41-ed9d-4ee5-95cc-c2699cab4abc",
    "X": 0,
    "Y": 0
}
```



***Responses:***


Status: http://localhost:8080/minesweeper/play | Code: 200



```js
{
    "game_id": "ec30fb3d-5f7e-4bc4-9724-f39f0498e45c",
    "token": "07d47f41-ed9d-4ee5-95cc-c2699cab4abc",
    "human_friendly_board": "[0][0][ ][ ][ ][ ][ ][ ][ ][ ]\n[0][0][ ][ ][ ][ ][ ][ ][ ][ ]\n[ ][ ][ ][ ][ ][ ][ ][ ][ ][ ]\n[ ][ ][ ][ ][ ][ ][ ][ ][ ][ ]\n[ ][ ][ ][ ][ ][ ][ ][ ][ ][ ]\n[ ][ ][ ][ ][ ][ ][ ][ ][ ][ ]\n[ ][ ][ ][ ][ ][ ][ ][ ][ ][ ]\n[ ][ ][ ][ ][ ][ ][ ][ ][ ][ ]\n[ ][ ][ ][ ][ ][ ][ ][ ][ ][ ]\n[ ][ ][ ][ ][ ][ ][ ][ ][ ][ ]\n",
    "board": [
        [
            {
                "clicked": true,
                "label": "[0]",
                "X": 0,
                "Y": 0
            },
            {
                "clicked": false,
                "label": "[0]",
                "X": 1,
                "Y": 0
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 2,
                "Y": 0
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 3,
                "Y": 0
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 4,
                "Y": 0
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 5,
                "Y": 0
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 6,
                "Y": 0
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 7,
                "Y": 0
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 8,
                "Y": 0
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 9,
                "Y": 0
            }
        ],
        [
            {
                "clicked": false,
                "label": "[0]",
                "X": 0,
                "Y": 1
            },
            {
                "clicked": false,
                "label": "[0]",
                "X": 1,
                "Y": 1
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 2,
                "Y": 1
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 3,
                "Y": 1
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 4,
                "Y": 1
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 5,
                "Y": 1
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 6,
                "Y": 1
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 7,
                "Y": 1
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 8,
                "Y": 1
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 9,
                "Y": 1
            }
        ],
        [
            {
                "clicked": false,
                "label": "[ ]",
                "X": 0,
                "Y": 2
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 1,
                "Y": 2
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 2,
                "Y": 2
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 3,
                "Y": 2
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 4,
                "Y": 2
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 5,
                "Y": 2
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 6,
                "Y": 2
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 7,
                "Y": 2
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 8,
                "Y": 2
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 9,
                "Y": 2
            }
        ],
        [
            {
                "clicked": false,
                "label": "[ ]",
                "X": 0,
                "Y": 3
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 1,
                "Y": 3
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 2,
                "Y": 3
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 3,
                "Y": 3
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 4,
                "Y": 3
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 5,
                "Y": 3
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 6,
                "Y": 3
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 7,
                "Y": 3
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 8,
                "Y": 3
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 9,
                "Y": 3
            }
        ],
        [
            {
                "clicked": false,
                "label": "[ ]",
                "X": 0,
                "Y": 4
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 1,
                "Y": 4
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 2,
                "Y": 4
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 3,
                "Y": 4
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 4,
                "Y": 4
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 5,
                "Y": 4
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 6,
                "Y": 4
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 7,
                "Y": 4
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 8,
                "Y": 4
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 9,
                "Y": 4
            }
        ],
        [
            {
                "clicked": false,
                "label": "[ ]",
                "X": 0,
                "Y": 5
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 1,
                "Y": 5
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 2,
                "Y": 5
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 3,
                "Y": 5
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 4,
                "Y": 5
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 5,
                "Y": 5
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 6,
                "Y": 5
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 7,
                "Y": 5
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 8,
                "Y": 5
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 9,
                "Y": 5
            }
        ],
        [
            {
                "clicked": false,
                "label": "[ ]",
                "X": 0,
                "Y": 6
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 1,
                "Y": 6
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 2,
                "Y": 6
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 3,
                "Y": 6
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 4,
                "Y": 6
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 5,
                "Y": 6
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 6,
                "Y": 6
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 7,
                "Y": 6
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 8,
                "Y": 6
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 9,
                "Y": 6
            }
        ],
        [
            {
                "clicked": false,
                "label": "[ ]",
                "X": 0,
                "Y": 7
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 1,
                "Y": 7
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 2,
                "Y": 7
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 3,
                "Y": 7
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 4,
                "Y": 7
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 5,
                "Y": 7
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 6,
                "Y": 7
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 7,
                "Y": 7
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 8,
                "Y": 7
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 9,
                "Y": 7
            }
        ],
        [
            {
                "clicked": false,
                "label": "[ ]",
                "X": 0,
                "Y": 8
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 1,
                "Y": 8
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 2,
                "Y": 8
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 3,
                "Y": 8
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 4,
                "Y": 8
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 5,
                "Y": 8
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 6,
                "Y": 8
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 7,
                "Y": 8
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 8,
                "Y": 8
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 9,
                "Y": 8
            }
        ],
        [
            {
                "clicked": false,
                "label": "[ ]",
                "X": 0,
                "Y": 9
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 1,
                "Y": 9
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 2,
                "Y": 9
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 3,
                "Y": 9
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 4,
                "Y": 9
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 5,
                "Y": 9
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 6,
                "Y": 9
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 7,
                "Y": 9
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 8,
                "Y": 9
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 9,
                "Y": 9
            }
        ]
    ],
    "game_setup": {
        "cols": 10,
        "rows": 10,
        "mines": 10
    },
    "game_status": "ACTIVE",
    "start_timestamp": 1573968107,
    "elapsed_time": 49400
}
```



### 4. http://localhost:8080/minesweeper/init



***Endpoint:***

```bash
Method: POST
Type: RAW
URL: http://localhost:8080/minesweeper/init
```


***Headers:***

| Key | Value | Description |
| --- | ------|-------------|
| Content-Type | application/json |  |



***Body:***

```js        
{
    "cols": 10,
    "rows": 10,
    "mines": 10
}
```



***Responses:***


Status: http://localhost:8080/minesweeper/init | Code: 201



```js
{
    "game_id": "ec30fb3d-5f7e-4bc4-9724-f39f0498e45c",
    "token": "07d47f41-ed9d-4ee5-95cc-c2699cab4abc",
    "human_friendly_board": "[ ][ ][ ][ ][ ][ ][ ][ ][ ][ ]\n[ ][ ][ ][ ][ ][ ][ ][ ][ ][ ]\n[ ][ ][ ][ ][ ][ ][ ][ ][ ][ ]\n[ ][ ][ ][ ][ ][ ][ ][ ][ ][ ]\n[ ][ ][ ][ ][ ][ ][ ][ ][ ][ ]\n[ ][ ][ ][ ][ ][ ][ ][ ][ ][ ]\n[ ][ ][ ][ ][ ][ ][ ][ ][ ][ ]\n[ ][ ][ ][ ][ ][ ][ ][ ][ ][ ]\n[ ][ ][ ][ ][ ][ ][ ][ ][ ][ ]\n[ ][ ][ ][ ][ ][ ][ ][ ][ ][ ]\n",
    "board": [
        [
            {
                "clicked": false,
                "label": "[ ]",
                "X": 0,
                "Y": 0
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 1,
                "Y": 0
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 2,
                "Y": 0
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 3,
                "Y": 0
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 4,
                "Y": 0
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 5,
                "Y": 0
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 6,
                "Y": 0
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 7,
                "Y": 0
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 8,
                "Y": 0
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 9,
                "Y": 0
            }
        ],
        [
            {
                "clicked": false,
                "label": "[ ]",
                "X": 0,
                "Y": 1
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 1,
                "Y": 1
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 2,
                "Y": 1
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 3,
                "Y": 1
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 4,
                "Y": 1
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 5,
                "Y": 1
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 6,
                "Y": 1
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 7,
                "Y": 1
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 8,
                "Y": 1
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 9,
                "Y": 1
            }
        ],
        [
            {
                "clicked": false,
                "label": "[ ]",
                "X": 0,
                "Y": 2
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 1,
                "Y": 2
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 2,
                "Y": 2
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 3,
                "Y": 2
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 4,
                "Y": 2
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 5,
                "Y": 2
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 6,
                "Y": 2
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 7,
                "Y": 2
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 8,
                "Y": 2
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 9,
                "Y": 2
            }
        ],
        [
            {
                "clicked": false,
                "label": "[ ]",
                "X": 0,
                "Y": 3
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 1,
                "Y": 3
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 2,
                "Y": 3
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 3,
                "Y": 3
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 4,
                "Y": 3
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 5,
                "Y": 3
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 6,
                "Y": 3
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 7,
                "Y": 3
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 8,
                "Y": 3
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 9,
                "Y": 3
            }
        ],
        [
            {
                "clicked": false,
                "label": "[ ]",
                "X": 0,
                "Y": 4
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 1,
                "Y": 4
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 2,
                "Y": 4
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 3,
                "Y": 4
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 4,
                "Y": 4
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 5,
                "Y": 4
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 6,
                "Y": 4
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 7,
                "Y": 4
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 8,
                "Y": 4
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 9,
                "Y": 4
            }
        ],
        [
            {
                "clicked": false,
                "label": "[ ]",
                "X": 0,
                "Y": 5
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 1,
                "Y": 5
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 2,
                "Y": 5
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 3,
                "Y": 5
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 4,
                "Y": 5
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 5,
                "Y": 5
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 6,
                "Y": 5
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 7,
                "Y": 5
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 8,
                "Y": 5
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 9,
                "Y": 5
            }
        ],
        [
            {
                "clicked": false,
                "label": "[ ]",
                "X": 0,
                "Y": 6
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 1,
                "Y": 6
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 2,
                "Y": 6
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 3,
                "Y": 6
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 4,
                "Y": 6
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 5,
                "Y": 6
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 6,
                "Y": 6
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 7,
                "Y": 6
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 8,
                "Y": 6
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 9,
                "Y": 6
            }
        ],
        [
            {
                "clicked": false,
                "label": "[ ]",
                "X": 0,
                "Y": 7
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 1,
                "Y": 7
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 2,
                "Y": 7
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 3,
                "Y": 7
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 4,
                "Y": 7
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 5,
                "Y": 7
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 6,
                "Y": 7
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 7,
                "Y": 7
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 8,
                "Y": 7
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 9,
                "Y": 7
            }
        ],
        [
            {
                "clicked": false,
                "label": "[ ]",
                "X": 0,
                "Y": 8
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 1,
                "Y": 8
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 2,
                "Y": 8
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 3,
                "Y": 8
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 4,
                "Y": 8
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 5,
                "Y": 8
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 6,
                "Y": 8
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 7,
                "Y": 8
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 8,
                "Y": 8
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 9,
                "Y": 8
            }
        ],
        [
            {
                "clicked": false,
                "label": "[ ]",
                "X": 0,
                "Y": 9
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 1,
                "Y": 9
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 2,
                "Y": 9
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 3,
                "Y": 9
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 4,
                "Y": 9
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 5,
                "Y": 9
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 6,
                "Y": 9
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 7,
                "Y": 9
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 8,
                "Y": 9
            },
            {
                "clicked": false,
                "label": "[ ]",
                "X": 9,
                "Y": 9
            }
        ]
    ],
    "game_setup": {
        "cols": 10,
        "rows": 10,
        "mines": 10
    },
    "game_status": "ACTIVE",
    "start_timestamp": 1573968107,
    "elapsed_time": 0
}
```



### 5. http://localhost:8080/ping



***Endpoint:***

```bash
Method: GET
Type: 
URL: http://localhost:8080/ping
```



***Responses:***


Status: http://localhost:8080/ping | Code: 200



```js
pong
```



---
[Back to top](#minesweeper)
> Made with &#9829; by guillermodoghel
