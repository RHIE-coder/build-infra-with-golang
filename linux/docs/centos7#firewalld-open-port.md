# CentOS 7 Firewalld 포트 열기

## PATH : `/etc/firewalld/zones/public.xml`

```xml
<zone>
  <short>Public</short>
  <description>For use in public areas. You do not trust the other computers on networks to not harm your computer. Only selected incoming connections are accepted.</description>
  <service name="ssh"/>
  <service name="dhcpv6-client"/>
  <port protocol="tcp" port="8080"/>
  <port protocol="tcp" port="80"/>
</zone>
```