# 커맨드 위치

get-command

# ExecutionPolicy

PowerShell(관리자 모드)

```cmd
Get-ExecutionPolicy                # Restricted
Get-ExecutionPolicy -List
#################################

        Scope ExecutionPolicy
        ----- ---------------
MachinePolicy       Undefined 모든 사용자정책을 그룹정책에서 설정한다.
   UserPolicy       Undefined 현재 사용자정책을 그룹정책에서 설정한다.
      Process       Undefined 현재 파워셸 세션에서만 유효하게 설정한다.
  CurrentUser       Undefined 현재 사용자에게만 유효하게 설정한다.
 LocalMachine       Undefined 기본 스코프로 모든 유저에게 영향을 끼친다.

#################################


Set-ExecutionPolicy RemoteSigned
Set-ExecutionPolicy -scope CurrentUser Unrestricted