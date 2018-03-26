# GET /health
display health status of the application

# GET /startstream
start streaming video with given parameters.
maximum 2 videos can be streamed at the same time.

## supported paramteres:
+ site (required): valid options: `cmore.se`, `cmore.dk`, `cmore.no`

+ user_id (required): id of user

+ video_id (required): id of the video to play

## examples
+ start streaming a video

request:
```
GET /startstream?site=cmore.se&video_id=3939304&user_id=1231
```
response:
```
{
  "VideoID": "3939304",
  "Status": "started",
  "Message": "video with id 3939304 found: Malou efter tio:"
}
```

+ invalid video id

request:
```
GET /startstream?site=cmore.se&video_id=1234567&user_id=1231
```
response:
```
{
  "VideoID": "1234567",
  "Status": "error",
  "Message": "video with id 1234567 not found"
}
```

+ breach limit of 2 video streams

request:
```
GET /startstream?site=cmore.se&video_id=3939304&user_id=1231
```
response:
```
{
  "VideoID": "3949520",
  "Status": "error",
  "Message": "users can only watch 2 video streams at the same time: map[3949520:true 3949522:true]"
}
```

# GET /stopstream
stop streaming video with given parameters.

## supported paramteres:
+ site (required): valid options: `cmore.se`, `cmore.dk`, `cmore.no`

+ user_id (required): id of user

+ video_id (required): id of the video to play

## examples
+ stop streaming a video

request:
```
GET /stopstream?site=cmore.se&video_id=3939304&user_id=1231
```
response:
```
{
  "VideoID": "3939304",
  "Status": "stopped",
  "Message": "video with id 3939304 found: Malou efter tio:"
}
```

+ video is not playing

request:
```
GET /stopstream?site=cmore.se&video_id=3939304&user_id=1231
```
response:
```
{
  "VideoID": "3939304",
  "Status": "error",
  "Message": "you're not watching video 3939304"
}
```
