https://learn.microsoft.com/ko-kr/sharepoint/troubleshoot/sync/delete-onedrive-synced-file-error

```
chkdsk C: /R /F
<재부팅>
해결! 이제 지울 수 있음.
```

 - 아래는 다른 옵션이지만 해결되지 않음
```
$ sfc /scannow
$ DISM /Online /Cleanup-Image /ScanHealth
$ DISM.exe /Online /Cleanup-image /Restorehealth
```