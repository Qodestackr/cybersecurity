# Detection Engineering Cheat Sheet

### Yara Rules

> Example

```yara
rule silent_banker : banker
{
    meta:
        description = "This is just an example"
        threat_level = 3
        in_the_wild = true
    strings:
        $a = {6A 40 68 00 30 00 00 6A 14 8D 91}
        $b = {8D 4D B0 2B C1 83 C0 27 99 6A 4E 59 F7 F9}
        $c = "UVODFRYSIHLNWPEJXQZAKCBGMT"
    condition:
        1 of ($*)
}
```

### Malware Analysis Tools

- [Process Monitor](https://docs.microsoft.com/en-us/sysinternals/downloads/procmon)
- [Process Explorer](https://docs.microsoft.com/en-us/sysinternals/downloads/process-explorer)
- [Process Hacker](https://processhacker.sourceforge.io/)
- [Autoruns](https://docs.microsoft.com/en-us/sysinternals/downloads/autoruns)
- [TCPView](https://docs.microsoft.com/en-us/sysinternals/downloads/tcpview)
- [Strings](https://docs.microsoft.com/en-us/sysinternals/downloads/strings)
- [CFF Explorer](https://ntcore.com/?page_id=388)
- 