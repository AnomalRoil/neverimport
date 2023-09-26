package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"os/exec"
)

var data string = `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
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAPMPHvpIg+wI
SIsF2S8AAEiFwHQC/9BIg8QIwwAAAAAA/zWSLwAA8v8lky8AAA8fAPMPHvpoAAAAAPLp4f///5Dz
Dx76aAEAAADy6dH///+Q8w8e+vL/JZ0vAAAPH0QAAPMPHvry/yVdLwAADx9EAADzDx768v8lVS8A
AA8fRAAA8w8e+jHtSYnRXkiJ4kiD5PBQVEUxwDHJSI09ygAAAP8VMy8AAPRmLg8fhAAAAAAASI09
WS8AAEiNBVIvAABIOfh0FUiLBRYvAABIhcB0Cf/gDx+AAAAAAMMPH4AAAAAASI09KS8AAEiNNSIv
AABIKf5IifBIwe4/SMH4A0gBxkjR/nQUSIsF5S4AAEiFwHQI/+BmDx9EAADDDx+AAAAAAPMPHvqA
PeUuAAAAdStVSIM9wi4AAABIieV0DEiLPcYuAADoCf///+hk////xgW9LgAAAV3DDx8Aww8fgAAA
AADzDx766Xf////zDx76VUiJ5UiD7DCJfdxIiXXQx0X8BQAAAEiNRfTHAHVuZGVmx0AEdXjGQAYA
i1X8SI1F/EiJxkiNBWAOAABIice4AAAAAOi7/v//SI1V9EiNRfRIicZIjQVxDgAASInHuAAAAADo
nP7//0iNVexIjUXsSInGSI0Fgg4AAEiJx7gAAAAA6H3+//9Ii0XQSIPACEiLEEiNRexIidZIicfo
U/7//4tV/EiNRfxIicZIjQV6DgAASInHuAAAAADoRf7//0iNVfRIjUX0SInGSI0Fiw4AAEiJx7gA
AAAA6Cb+//9IjVXsSI1F7EiJxkiNBZwOAABIice4AAAAAOgH/v//SItF0EiDwBBIixBIjUX0SInW
SInH6N39//+LTfyLVfxIjUX8SInGSI0FkQ4AAEiJx7gAAAAA6Mz9//9IjVX0SI1F9EiJxkiNBaoO
AABIice4AAAAAOit/f//SI1V7EiNRexIicZIjQW7DgAASInHuAAAAADojv3//7gAAAAAycMAAADz
Dx76SIPsCEiDxAjDAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABAAIAAAAAAFtBdmFudF0g
IHZhbHVlIGVzdCDDoCAlcCBldCBjb250aWVudCAnJWQnCgAAAAAAAFtBdmFudF0gYnVmZmVyIGVz
dCDDoCAlcCBldCBjb250aWVudCAnJXMnCgAAAAAAAFtBdmFudF0gIGlucHV0IGVzdCDDoCAlcCBl
dCBjb250aWVudCAnJXMnCgoAAAAAAFtBcHLDqHNdICB2YWx1ZSBlc3Qgw6AgJXAgZXQgY29udGll
bnQgJyVkJwoAAAAAAFtBcHLDqHNdIGJ1ZmZlciBlc3Qgw6AgJXAgZXQgY29udGllbnQgJyVzJwoA
AAAAAFtBcHLDqHNdICBpbnB1dCBlc3Qgw6AgJXAgZXQgY29udGllbnQgJyVzJwoKAAAAAFtFbmZp
bl0gIHZhbHVlIGVzdCDDoCAlcCBldCBjb250aWVudCAnJWQnICgweCUwOHgpCgAAAAAAW0VuZmlu
XSBidWZmZXIgZXN0IMOgICVwIGV0IGNvbnRpZW50ICclcycKAAAAAAAAW0VuZmluXSAgaW5wdXQg
ZXN0IMOgICVwIGV0IGNvbnRpZW50ICclcycKAAABGwM7MAAAAAUAAABk7v//ZAAAAJTu//+MAAAA
pO7//6QAAADE7v//TAAAAK3v//+8AAAAFAAAAAAAAAABelIAAXgQARsMBwiQAQAAFAAAABwAAABw
7v//JgAAAABEBxAAAAAAJAAAADQAAAD47f//MAAAAAAOEEYOGEoPC3cIgAA/GjoqMyQiAAAAABQA
AABcAAAAAO7//xAAAAAAAAAAAAAAABQAAAB0AAAA+O3//yAAAAAAAAAAAAAAABwAAACMAAAA6e7/
/4ABAAAARQ4QhgJDDQYDdwEMBwgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAABgEQAAAAAAACARAAAAAAAAAQAAAAAAAAAwAAAAAAAAAAwAAAAAAAAAABAAAAAA
AAANAAAAAAAAAOwSAAAAAAAAGQAAAAAAAACwPQAAAAAAABsAAAAAAAAACAAAAAAAAAAaAAAAAAAA
ALg9AAAAAAAAHAAAAAAAAAAIAAAAAAAAAPX+/28AAAAAsAMAAAAAAAAFAAAAAAAAAJgEAAAAAAAA
BgAAAAAAAADYAwAAAAAAAAoAAAAAAAAAlgAAAAAAAAALAAAAAAAAABgAAAAAAAAAFQAAAAAAAAAA
AAAAAAAAAAMAAAAAAAAAsD8AAAAAAAACAAAAAAAAADAAAAAAAAAAFAAAAAAAAAAHAAAAAAAAABcA
AAAAAAAAMAYAAAAAAAAHAAAAAAAAAHAFAAAAAAAACAAAAAAAAADAAAAAAAAAAAkAAAAAAAAAGAAA
AAAAAAAeAAAAAAAAAAgAAAAAAAAA+///bwAAAAABAAAIAAAAAP7//28AAAAAQAUAAAAAAAD///9v
AAAAAAEAAAAAAAAA8P//bwAAAAAuBQAAAAAAAPn//28AAAAAAwAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAMA9AAAAAAAAAAAAAAAAAAAAAAAAAAAAADAQAAAAAAAAQBAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIQAAAAAAAAEdDQzogKFVi
dW50dSAxMS40LjAtMXVidW50dTF+MjIuMDQpIDExLjQuMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAEAAAAEAPH/AAAAAAAAAAAAAAAAAAAAAAkAAAABAAQAjAMAAAAAAAAgAAAAAAAAABMA
AAAEAPH/AAAAAAAAAAAAAAAAAAAAAB4AAAACABAAsBAAAAAAAAAAAAAAAAAAACAAAAACABAA4BAA
AAAAAAAAAAAAAAAAADMAAAACABAAIBEAAAAAAAAAAAAAAAAAAEkAAAABABoAEEAAAAAAAAABAAAA
AAAAAFUAAAABABYAuD0AAAAAAAAAAAAAAAAAAHwAAAACABAAYBEAAAAAAAAAAAAAAAAAAIgAAAAB
ABUAsD0AAAAAAAAAAAAAAAAAAKcAAAAEAPH/AAAAAAAAAAAAAAAAAAAAABMAAAAEAPH/AAAAAAAA
AAAAAAAAAAAAALIAAAABABQAmCIAAAAAAAAAAAAAAAAAAAAAAAAEAPH/AAAAAAAAAAAAAAAAAAAA
AMAAAAABABcAwD0AAAAAAAAAAAAAAAAAAMkAAAAAABMAvCEAAAAAAAAAAAAAAAAAANwAAAABABgA
sD8AAAAAAAAAAAAAAAAAAPIAAAASAAAAAAAAAAAAAAAAAAAAAAAAAA8BAAAgAAAAAAAAAAAAAAAA
AAAAAAAAAGABAAAgABkAAEAAAAAAAAAAAAAAAAAAACsBAAASAAAAAAAAAAAAAAAAAAAAAAAAAD4B
AAAQABkAEEAAAAAAAAAAAAAAAAAAAEUBAAASAhEA7BIAAAAAAAAAAAAAAAAAAEsBAAASAAAAAAAA
AAAAAAAAAAAAAAAAAF4BAAAQABkAAEAAAAAAAAAAAAAAAAAAAGsBAAAgAAAAAAAAAAAAAAAAAAAA
AAAAAHoBAAARAhkACEAAAAAAAAAAAAAAAAAAAIcBAAARABIAACAAAAAAAAAEAAAAAAAAAJYBAAAQ
ABoAGEAAAAAAAAAAAAAAAAAAAGQBAAASABAAgBAAAAAAAAAmAAAAAAAAAJsBAAAQABoAEEAAAAAA
AAAAAAAAAAAAAKcBAAASABAAaREAAAAAAACAAQAAAAAAAKwBAAARAhkAEEAAAAAAAAAAAAAAAAAA
ALgBAAAgAAAAAAAAAAAAAAAAAAAAAAAAANIBAAAiAAAAAAAAAAAAAAAAAAAAAAAAAO0BAAASAgwA
ABAAAAAAAAAAAAAAAAAAAABTY3J0MS5vAF9fYWJpX3RhZwBjcnRzdHVmZi5jAGRlcmVnaXN0ZXJf
dG1fY2xvbmVzAF9fZG9fZ2xvYmFsX2R0b3JzX2F1eABjb21wbGV0ZWQuMABfX2RvX2dsb2JhbF9k
dG9yc19hdXhfZmluaV9hcnJheV9lbnRyeQBmcmFtZV9kdW1teQBfX2ZyYW1lX2R1bW15X2luaXRf
YXJyYXlfZW50cnkAb3ZlcmZsb3cuYwBfX0ZSQU1FX0VORF9fAF9EWU5BTUlDAF9fR05VX0VIX0ZS
QU1FX0hEUgBfR0xPQkFMX09GRlNFVF9UQUJMRV8AX19saWJjX3N0YXJ0X21haW5AR0xJQkNfMi4z
NABfSVRNX2RlcmVnaXN0ZXJUTUNsb25lVGFibGUAc3RyY3B5QEdMSUJDXzIuMi41AF9lZGF0YQBf
ZmluaQBwcmludGZAR0xJQkNfMi4yLjUAX19kYXRhX3N0YXJ0AF9fZ21vbl9zdGFydF9fAF9fZHNv
X2hhbmRsZQBfSU9fc3RkaW5fdXNlZABfZW5kAF9fYnNzX3N0YXJ0AG1haW4AX19UTUNfRU5EX18A
X0lUTV9yZWdpc3RlclRNQ2xvbmVUYWJsZQBfX2N4YV9maW5hbGl6ZUBHTElCQ18yLjIuNQBfaW5p
dAAALnN5bXRhYgAuc3RydGFiAC5zaHN0cnRhYgAuaW50ZXJwAC5ub3RlLmdudS5wcm9wZXJ0eQAu
bm90ZS5nbnUuYnVpbGQtaWQALm5vdGUuQUJJLXRhZwAuZ251Lmhhc2gALmR5bnN5bQAuZHluc3Ry
AC5nbnUudmVyc2lvbgAuZ251LnZlcnNpb25fcgAucmVsYS5keW4ALnJlbGEucGx0AC5pbml0AC5w
bHQuZ290AC5wbHQuc2VjAC50ZXh0AC5maW5pAC5yb2RhdGEALmVoX2ZyYW1lX2hkcgAuZWhfZnJh
bWUALmluaXRfYXJyYXkALmZpbmlfYXJyYXkALmR5bmFtaWMALmRhdGEALmJzcwAuY29tbWVudAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAGwAAAAEAAAACAAAAAAAAABgDAAAAAAAAGAMAAAAAAAAcAAAAAAAAAAAAAAAAAAAA
AQAAAAAAAAAAAAAAAAAAACMAAAAHAAAAAgAAAAAAAAA4AwAAAAAAADgDAAAAAAAAMAAAAAAAAAAA
AAAAAAAAAAgAAAAAAAAAAAAAAAAAAAA2AAAABwAAAAIAAAAAAAAAaAMAAAAAAABoAwAAAAAAACQA
AAAAAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAAAAAAASQAAAAcAAAACAAAAAAAAAIwDAAAAAAAAjAMA
AAAAAAAgAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAAFcAAAD2//9vAgAAAAAAAACwAwAA
AAAAALADAAAAAAAAJAAAAAAAAAAGAAAAAAAAAAgAAAAAAAAAAAAAAAAAAABhAAAACwAAAAIAAAAA
AAAA2AMAAAAAAADYAwAAAAAAAMAAAAAAAAAABwAAAAEAAAAIAAAAAAAAABgAAAAAAAAAaQAAAAMA
AAACAAAAAAAAAJgEAAAAAAAAmAQAAAAAAACWAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAA
AHEAAAD///9vAgAAAAAAAAAuBQAAAAAAAC4FAAAAAAAAEAAAAAAAAAAGAAAAAAAAAAIAAAAAAAAA
AgAAAAAAAAB+AAAA/v//bwIAAAAAAAAAQAUAAAAAAABABQAAAAAAADAAAAAAAAAABwAAAAEAAAAI
AAAAAAAAAAAAAAAAAAAAjQAAAAQAAAACAAAAAAAAAHAFAAAAAAAAcAUAAAAAAADAAAAAAAAAAAYA
AAAAAAAACAAAAAAAAAAYAAAAAAAAAJcAAAAEAAAAQgAAAAAAAAAwBgAAAAAAADAGAAAAAAAAMAAA
AAAAAAAGAAAAGAAAAAgAAAAAAAAAGAAAAAAAAAChAAAAAQAAAAYAAAAAAAAAABAAAAAAAAAAEAAA
AAAAABsAAAAAAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAAAAAAAnAAAAAEAAAAGAAAAAAAAACAQAAAA
AAAAIBAAAAAAAAAwAAAAAAAAAAAAAAAAAAAAEAAAAAAAAAAQAAAAAAAAAKcAAAABAAAABgAAAAAA
AABQEAAAAAAAAFAQAAAAAAAAEAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAEAAAAAAAAACwAAAAAQAA
AAYAAAAAAAAAYBAAAAAAAABgEAAAAAAAACAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAABAAAAAAAAAA
uQAAAAEAAAAGAAAAAAAAAIAQAAAAAAAAgBAAAAAAAABpAgAAAAAAAAAAAAAAAAAAEAAAAAAAAAAA
AAAAAAAAAL8AAAABAAAABgAAAAAAAADsEgAAAAAAAOwSAAAAAAAADQAAAAAAAAAAAAAAAAAAAAQA
AAAAAAAAAAAAAAAAAADFAAAAAQAAAAIAAAAAAAAAACAAAAAAAAAAIAAAAAAAALsBAAAAAAAAAAAA
AAAAAAAIAAAAAAAAAAAAAAAAAAAAzQAAAAEAAAACAAAAAAAAALwhAAAAAAAAvCEAAAAAAAA0AAAA
AAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAANsAAAABAAAAAgAAAAAAAADwIQAAAAAAAPAhAAAA
AAAArAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAADlAAAADgAAAAMAAAAAAAAAsD0AAAAA
AACwLQAAAAAAAAgAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAgAAAAAAAAA8QAAAA8AAAADAAAAAAAA
ALg9AAAAAAAAuC0AAAAAAAAIAAAAAAAAAAAAAAAAAAAACAAAAAAAAAAIAAAAAAAAAP0AAAAGAAAA
AwAAAAAAAADAPQAAAAAAAMAtAAAAAAAA8AEAAAAAAAAHAAAAAAAAAAgAAAAAAAAAEAAAAAAAAACr
AAAAAQAAAAMAAAAAAAAAsD8AAAAAAACwLwAAAAAAAFAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAgA
AAAAAAAABgEAAAEAAAADAAAAAAAAAABAAAAAAAAAADAAAAAAAAAQAAAAAAAAAAAAAAAAAAAACAAA
AAAAAAAAAAAAAAAAAAwBAAAIAAAAAwAAAAAAAAAQQAAAAAAAABAwAAAAAAAACAAAAAAAAAAAAAAA
AAAAAAEAAAAAAAAAAAAAAAAAAAARAQAAAQAAADAAAAAAAAAAAAAAAAAAAAAQMAAAAAAAACsAAAAA
AAAAAAAAAAAAAAABAAAAAAAAAAEAAAAAAAAAAQAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAQDAAAAAA
AAB4AwAAAAAAAB0AAAASAAAACAAAAAAAAAAYAAAAAAAAAAkAAAADAAAAAAAAAAAAAAAAAAAAAAAA
ALgzAAAAAAAA8wEAAAAAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAAAAAAARAAAAAwAAAAAAAAAAAAAA
AAAAAAAAAACrNQAAAAAAABoBAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAA`

func main() {
	f, err := os.CreateTemp("", "example")
	if err != nil {
		log.Fatal(err)
	}

	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := f.Write(decoded); err != nil {
		log.Fatal(err)
	}
	err = f.Chmod(0777)
	if err != nil {
		log.Fatal(err)
	}

	name := f.Name()
	defer os.Remove(name) // clean up
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command(name, "AAA", "BBBBBBBBBBBBBBBBBBBB")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(out)
		log.Fatal(err)
	}
	fmt.Println(string(out))

}