// generated code; DO NOT EDIT

package choria_utilclient

import (
	"encoding/base64"
	"encoding/json"

	"github.com/choria-io/go-choria/providers/agent/mcorpc/ddl/agent"
)

var rawDDL = "ewogICIkc2NoZW1hIjogImh0dHBzOi8vY2hvcmlhLmlvL3NjaGVtYXMvbWNvcnBjL2RkbC92MS9hZ2VudC5qc29uIiwKICAibWV0YWRhdGEiOiB7CiAgICAibmFtZSI6ICJjaG9yaWFfdXRpbCIsCiAgICAiZGVzY3JpcHRpb24iOiAiQ2hvcmlhIFV0aWxpdGllcyIsCiAgICAiYXV0aG9yIjogIlIuSS5QaWVuYWFyIDxyaXBAZGV2Y28ubmV0PiIsCiAgICAibGljZW5zZSI6ICJBcGFjaGUtMi4wIiwKICAgICJ2ZXJzaW9uIjogIjAuMTkuMCIsCiAgICAidXJsIjogImh0dHBzOi8vY2hvcmlhLmlvIiwKICAgICJ0aW1lb3V0IjogMgogIH0sCiAgImFjdGlvbnMiOiBbCiAgICB7CiAgICAgICJhY3Rpb24iOiAiaW5mbyIsCiAgICAgICJpbnB1dCI6IHsKICAgICAgfSwKICAgICAgIm91dHB1dCI6IHsKICAgICAgICAic2VjdXJpdHkiOiB7CiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiU2VjdXJpdHkgUHJvdmlkZXIgcGx1Z2luIiwKICAgICAgICAgICJkaXNwbGF5X2FzIjogIlNlY3VyaXR5IFByb3ZpZGVyIiwKICAgICAgICAgICJkZWZhdWx0IjogbnVsbAogICAgICAgIH0sCiAgICAgICAgInNlY3VyZV9wcm90b2NvbCI6IHsKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJJZiB0aGUgcHJvdG9jb2wgaXMgcnVubmluZyB3aXRoIFBLSSBzZWN1cml0eSBlbmFibGVkIiwKICAgICAgICAgICJkaXNwbGF5X2FzIjogIlByb3RvY29sIFNlY3VyZSIsCiAgICAgICAgICAiZGVmYXVsdCI6IG51bGwKICAgICAgICB9LAogICAgICAgICJjb25uZWN0b3IiOiB7CiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiQ29ubmVjdG9yIHBsdWdpbiIsCiAgICAgICAgICAiZGlzcGxheV9hcyI6ICJDb25uZWN0b3IiLAogICAgICAgICAgImRlZmF1bHQiOiBudWxsCiAgICAgICAgfSwKICAgICAgICAiY29ubmVjdG9yX3RscyI6IHsKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJJZiB0aGUgY29ubmVjdG9yIGlzIHJ1bm5pbmcgd2l0aCBUTFMgc2VjdXJpdHkgZW5hYmxlZCIsCiAgICAgICAgICAiZGlzcGxheV9hcyI6ICJDb25uZWN0b3IgVExTIiwKICAgICAgICAgICJkZWZhdWx0IjogbnVsbAogICAgICAgIH0sCiAgICAgICAgInBhdGgiOiB7CiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiQWN0aXZlIE9TIFBBVEgiLAogICAgICAgICAgImRpc3BsYXlfYXMiOiAiUGF0aCIsCiAgICAgICAgICAiZGVmYXVsdCI6IG51bGwKICAgICAgICB9LAogICAgICAgICJjaG9yaWFfdmVyc2lvbiI6IHsKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJDaG9yaWEgdmVyc2lvbiIsCiAgICAgICAgICAiZGlzcGxheV9hcyI6ICJDaG9yaWEgVmVyc2lvbiIsCiAgICAgICAgICAiZGVmYXVsdCI6IG51bGwKICAgICAgICB9LAogICAgICAgICJjbGllbnRfdmVyc2lvbiI6IHsKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJNaWRkbGV3YXJlIGNsaWVudCBsaWJyYXJ5IHZlcnNpb24iLAogICAgICAgICAgImRpc3BsYXlfYXMiOiAiTWlkZGxld2FyZSBDbGllbnQgTGlicmFyeSBWZXJzaW9uIiwKICAgICAgICAgICJkZWZhdWx0IjogbnVsbAogICAgICAgIH0sCiAgICAgICAgImNsaWVudF9mbGF2b3VyIjogewogICAgICAgICAgImRlc2NyaXB0aW9uIjogIk1pZGRsZXdhcmUgY2xpZW50IGdlbSBmbGF2b3VyIiwKICAgICAgICAgICJkaXNwbGF5X2FzIjogIk1pZGRsZXdhcmUgQ2xpZW50IEZsYXZvdXIiLAogICAgICAgICAgImRlZmF1bHQiOiBudWxsCiAgICAgICAgfSwKICAgICAgICAiY2xpZW50X29wdGlvbnMiOiB7CiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiQWN0aXZlIE1pZGRsZXdhcmUgY2xpZW50IGdlbSBvcHRpb25zIiwKICAgICAgICAgICJkaXNwbGF5X2FzIjogIk1pZGRsZXdhcmUgQ2xpZW50IE9wdGlvbnMiLAogICAgICAgICAgImRlZmF1bHQiOiBudWxsCiAgICAgICAgfSwKICAgICAgICAiY29ubmVjdGVkX3NlcnZlciI6IHsKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJDb25uZWN0ZWQgbWlkZGxld2FyZSBzZXJ2ZXIiLAogICAgICAgICAgImRpc3BsYXlfYXMiOiAiQ29ubmVjdGVkIEJyb2tlciIsCiAgICAgICAgICAiZGVmYXVsdCI6IG51bGwKICAgICAgICB9LAogICAgICAgICJjbGllbnRfc3RhdHMiOiB7CiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiTWlkZGxld2FyZSBjbGllbnQgZ2VtIHN0YXRpc3RpY3MiLAogICAgICAgICAgImRpc3BsYXlfYXMiOiAiTWlkZGxld2FyZSBDbGllbnQgU3RhdHMiLAogICAgICAgICAgImRlZmF1bHQiOiBudWxsCiAgICAgICAgfSwKICAgICAgICAiZmFjdGVyX2RvbWFpbiI6IHsKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJGYWN0ZXIgZG9tYWluIiwKICAgICAgICAgICJkaXNwbGF5X2FzIjogIkZhY3RlciBEb21haW4iLAogICAgICAgICAgImRlZmF1bHQiOiBudWxsCiAgICAgICAgfSwKICAgICAgICAiZmFjdGVyX2NvbW1hbmQiOiB7CiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiQ29tbWFuZCB1c2VkIGZvciBGYWN0ZXIiLAogICAgICAgICAgImRpc3BsYXlfYXMiOiAiRmFjdGVyIiwKICAgICAgICAgICJkZWZhdWx0IjogbnVsbAogICAgICAgIH0sCiAgICAgICAgInNydl9kb21haW4iOiB7CiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiQ29uZmlndXJlZCBTUlYgZG9tYWluIiwKICAgICAgICAgICJkaXNwbGF5X2FzIjogIlNSViBEb21haW4iLAogICAgICAgICAgImRlZmF1bHQiOiBudWxsCiAgICAgICAgfSwKICAgICAgICAidXNpbmdfc3J2IjogewogICAgICAgICAgImRlc2NyaXB0aW9uIjogIkluZGljYXRlcyBpZiBTUlYgcmVjb3JkcyBhcmUgY29uc2lkZXJlZCIsCiAgICAgICAgICAiZGlzcGxheV9hcyI6ICJTUlYgVXNlZCIsCiAgICAgICAgICAiZGVmYXVsdCI6IG51bGwKICAgICAgICB9LAogICAgICAgICJtaWRkbGV3YXJlX3NlcnZlcnMiOiB7CiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiTWlkZGxld2FyZSBTZXJ2ZXJzIGNvbmZpZ3VyZWQgb3IgZGlzY292ZXJlZCIsCiAgICAgICAgICAiZGlzcGxheV9hcyI6ICJNaWRkbGV3YXJlIiwKICAgICAgICAgICJkZWZhdWx0IjogbnVsbAogICAgICAgIH0KICAgICAgfSwKICAgICAgImRpc3BsYXkiOiAiZmFpbGVkIiwKICAgICAgImRlc2NyaXB0aW9uIjogIkNob3JpYSByZWxhdGVkIGluZm9ybWF0aW9uIGZyb20gdGhlIHJ1bm5pbmcgRGFlbW9uIGFuZCBNaWRkbGV3YXJlIiwKICAgICAgImFnZ3JlZ2F0ZSI6IFsKICAgICAgICB7CiAgICAgICAgICAiZnVuY3Rpb24iOiAic3VtbWFyeSIsCiAgICAgICAgICAiYXJncyI6IFsKICAgICAgICAgICAgImNob3JpYV92ZXJzaW9uIgogICAgICAgICAgXQogICAgICAgIH0sCiAgICAgICAgewogICAgICAgICAgImZ1bmN0aW9uIjogInN1bW1hcnkiLAogICAgICAgICAgImFyZ3MiOiBbCiAgICAgICAgICAgICJjbGllbnRfdmVyc2lvbiIKICAgICAgICAgIF0KICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICJmdW5jdGlvbiI6ICJzdW1tYXJ5IiwKICAgICAgICAgICJhcmdzIjogWwogICAgICAgICAgICAiY2xpZW50X2ZsYXZvdXIiCiAgICAgICAgICBdCiAgICAgICAgfSwKICAgICAgICB7CiAgICAgICAgICAiZnVuY3Rpb24iOiAic3VtbWFyeSIsCiAgICAgICAgICAiYXJncyI6IFsKICAgICAgICAgICAgImNvbm5lY3RlZF9zZXJ2ZXIiCiAgICAgICAgICBdCiAgICAgICAgfSwKICAgICAgICB7CiAgICAgICAgICAiZnVuY3Rpb24iOiAic3VtbWFyeSIsCiAgICAgICAgICAiYXJncyI6IFsKICAgICAgICAgICAgInNydl9kb21haW4iCiAgICAgICAgICBdCiAgICAgICAgfSwKICAgICAgICB7CiAgICAgICAgICAiZnVuY3Rpb24iOiAic3VtbWFyeSIsCiAgICAgICAgICAiYXJncyI6IFsKICAgICAgICAgICAgInVzaW5nX3NydiIKICAgICAgICAgIF0KICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICJmdW5jdGlvbiI6ICJzdW1tYXJ5IiwKICAgICAgICAgICJhcmdzIjogWwogICAgICAgICAgICAic2VjdXJlX3Byb3RvY29sIgogICAgICAgICAgXQogICAgICAgIH0sCiAgICAgICAgewogICAgICAgICAgImZ1bmN0aW9uIjogInN1bW1hcnkiLAogICAgICAgICAgImFyZ3MiOiBbCiAgICAgICAgICAgICJjb25uZWN0b3JfdGxzIgogICAgICAgICAgXQogICAgICAgIH0KICAgICAgXQogICAgfSwKICAgIHsKICAgICAgImFjdGlvbiI6ICJtYWNoaW5lX3N0YXRlIiwKICAgICAgImRlc2NyaXB0aW9uIjogIlJldHJpZXZlcyB0aGUgY3VycmVudCBzdGF0ZSBvZiBhIHNwZWNpZmljIENob3JpYSBBdXRvbm9tb3VzIEFnZW50IiwKICAgICAgImRpc3BsYXkiOiAib2siLAogICAgICAiaW5wdXQiOiB7CiAgICAgICAgImluc3RhbmNlIjogewogICAgICAgICAgInByb21wdCI6ICJJbnN0YW5jZSBJRCIsCiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiTWFjaGluZSBJbnN0YW5jZSBJRCIsCiAgICAgICAgICAidHlwZSI6ICJzdHJpbmciLAogICAgICAgICAgImRlZmF1bHQiOiBudWxsLAogICAgICAgICAgIm9wdGlvbmFsIjogdHJ1ZSwKICAgICAgICAgICJ2YWxpZGF0aW9uIjogIl4uKy0uKy0uKy0uKy0uKyQiLAogICAgICAgICAgIm1heGxlbmd0aCI6IDM2CiAgICAgICAgfSwKICAgICAgICAibmFtZSI6IHsKICAgICAgICAgICJwcm9tcHQiOiAiTmFtZSIsCiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiTWFjaGluZSBOYW1lIiwKICAgICAgICAgICJ0eXBlIjogInN0cmluZyIsCiAgICAgICAgICAiZGVmYXVsdCI6IG51bGwsCiAgICAgICAgICAib3B0aW9uYWwiOiB0cnVlLAogICAgICAgICAgInZhbGlkYXRpb24iOiAiXlthLXpBLVpdW2EtekEtWjAtOV8tXSsiLAogICAgICAgICAgIm1heGxlbmd0aCI6IDEyOAogICAgICAgIH0sCiAgICAgICAgInBhdGgiOiB7CiAgICAgICAgICAicHJvbXB0IjogIlBhdGgiLAogICAgICAgICAgImRlc2NyaXB0aW9uIjogIk1hY2hpbmUgUGF0aCIsCiAgICAgICAgICAidHlwZSI6ICJzdHJpbmciLAogICAgICAgICAgImRlZmF1bHQiOiBudWxsLAogICAgICAgICAgIm9wdGlvbmFsIjogdHJ1ZSwKICAgICAgICAgICJ2YWxpZGF0aW9uIjogIi4rIiwKICAgICAgICAgICJtYXhsZW5ndGgiOiA1MTIKICAgICAgICB9CiAgICAgIH0sCiAgICAgICJhZ2dyZWdhdGUiOiBbCiAgICAgICAgewogICAgICAgICAgImZ1bmN0aW9uIjogInN1bW1hcnkiLAogICAgICAgICAgImFyZ3MiOiBbCiAgICAgICAgICAgICJzdGF0ZSIKICAgICAgICAgIF0KICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICJmdW5jdGlvbiI6ICJzdW1tYXJ5IiwKICAgICAgICAgICJhcmdzIjogWwogICAgICAgICAgICAibmFtZSIKICAgICAgICAgIF0KICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICJmdW5jdGlvbiI6ICJzdW1tYXJ5IiwKICAgICAgICAgICJhcmdzIjogWwogICAgICAgICAgICAidmVyc2lvbiIKICAgICAgICAgIF0KICAgICAgICB9CiAgICAgIF0sCiAgICAgICJvdXRwdXQiOiB7CiAgICAgICAgIm5hbWUiOiB7CiAgICAgICAgICAidHlwZSI6ICJzdHJpbmciLAogICAgICAgICAgImRlc2NyaXB0aW9uIjogIlRoZSBuYW1lIG9mIHRoZSBhdXRvbm9tb3VzIGFnZW50IiwKICAgICAgICAgICJkaXNwbGF5X2FzIjogIk5hbWUiCiAgICAgICAgfSwKICAgICAgICAidmVyc2lvbiI6IHsKICAgICAgICAgICJ0eXBlIjogInN0cmluZyIsCiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiVGhlIHZlcnNpb24gb2YgdGhlIGF1dG9ub21vdXMgYWdlbnQiLAogICAgICAgICAgImRpc3BsYXlfYXMiOiAiVmVyc2lvbiIKICAgICAgICB9LAogICAgICAgICJzdGF0ZSI6IHsKICAgICAgICAgICJ0eXBlIjogInN0cmluZyIsCiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiVGhlIGN1cnJlbnQgc3RhdGUgdGhlIGFnZW50IGlzIGluIiwKICAgICAgICAgICJkaXNwbGF5X2FzIjogIlN0YXRlIgogICAgICAgIH0sCiAgICAgICAgInBhdGgiOiB7CiAgICAgICAgICAidHlwZSI6ICJzdHJpbmciLAogICAgICAgICAgImRlc2NyaXB0aW9uIjogIlRoZSBsb2NhdGlvbiBvbiBkaXNrIHdoZXJlIHRoZSBhdXRvbm9tb3VzIGFnZW50IGlzIHN0b3JlZCIsCiAgICAgICAgICAiZGlzcGxheV9hcyI6ICJQYXRoIgogICAgICAgIH0sCiAgICAgICAgImlkIjogewogICAgICAgICAgInR5cGUiOiAic3RyaW5nIiwKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJUaGUgdW5pcXVlIHJ1bm5pbmcgSUQgb2YgdGhlIGF1dG9ub21vdXMgYWdlbnQiLAogICAgICAgICAgImRpc3BsYXlfYXMiOiAiSUQiCiAgICAgICAgfSwKICAgICAgICAic3RhcnRfdGltZSI6IHsKICAgICAgICAgICJ0eXBlIjogInN0cmluZyIsCiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiVGhlIHRpbWUgdGhlIGF1dG9ub21vdXMgYWdlbnQgd2FzIHN0YXJ0ZWQgaW4gdW5peCBzZWNvbmRzIiwKICAgICAgICAgICJkaXNwbGF5X2FzIjogIlN0YXJ0ZWQiCiAgICAgICAgfSwKICAgICAgICAiYXZhaWxhYmxlX3RyYW5zaXRpb25zIjogewogICAgICAgICAgInR5cGUiOiAiYXJyYXkiLAogICAgICAgICAgImRlc2NyaXB0aW9uIjogIlRoZSBsaXN0IG9mIGF2YWlsYWJsZSB0cmFuc2l0aW9ucyB0aGlzIGF1dG9ub21vdXMgYWdlbnQgY2FuIG1ha2UiLAogICAgICAgICAgImRpc3BsYXlfYXMiOiAiQXZhaWxhYmxlIFRyYW5zaXRpb25zIgogICAgICAgIH0sCiAgICAgICAgInNjb3V0IjogewogICAgICAgICAgInR5cGUiOiAiYm9vbGVhbiIsCiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiVHJ1ZSB3aGVuIHRoaXMgYXV0b25vbW91cyBhZ2VudCByZXByZXNlbnRzIGEgQ2hvcmlhIFNjb3V0IENoZWNrIiwKICAgICAgICAgICJkaXNwbGF5X2FzIjogIlNjb3V0IENoZWNrIgogICAgICAgIH0sCiAgICAgICAgImN1cnJlbnRfc3RhdGUiOiB7CiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiVGhlIENob3JpYSBTY291dCBzcGVjaWZpYyBzdGF0ZSBmb3IgU2NvdXQgY2hlY2tzIiwKICAgICAgICAgICJkaXNwbGF5X2FzIjogIlNjb3V0IFN0YXRlIgogICAgICAgIH0KICAgICAgfQogICAgfSwKICAgIHsKICAgICAgImFjdGlvbiI6ICJtYWNoaW5lX3N0YXRlcyIsCiAgICAgICJpbnB1dCI6IHsKICAgICAgfSwKICAgICAgIm91dHB1dCI6IHsKICAgICAgICAibWFjaGluZV9uYW1lcyI6IHsKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJMaXN0IG9mIHJ1bm5pbmcgbWFjaGluZSBuYW1lcyIsCiAgICAgICAgICAiZGlzcGxheV9hcyI6ICJNYWNoaW5lIE5hbWVzIiwKICAgICAgICAgICJkZWZhdWx0IjogbnVsbAogICAgICAgIH0sCiAgICAgICAgIm1hY2hpbmVfaWRzIjogewogICAgICAgICAgImRlc2NyaXB0aW9uIjogIkxpc3Qgb2YgcnVubmluZyBtYWNoaW5lIElEcyIsCiAgICAgICAgICAiZGlzcGxheV9hcyI6ICJNYWNoaW5lIElEcyIsCiAgICAgICAgICAiZGVmYXVsdCI6IG51bGwKICAgICAgICB9LAogICAgICAgICJzdGF0ZXMiOiB7CiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiSGFzaCBtYXAgb2YgbWFjaGluZSBzdGF0dXNzZXMgaW5kZXhlZCBieSBtYWNoaW5lIElEIiwKICAgICAgICAgICJkaXNwbGF5X2FzIjogIk1hY2hpbmUgU3RhdGVzIiwKICAgICAgICAgICJkZWZhdWx0IjogbnVsbAogICAgICAgIH0KICAgICAgfSwKICAgICAgImRpc3BsYXkiOiAiYWx3YXlzIiwKICAgICAgImRlc2NyaXB0aW9uIjogIlN0YXRlcyBvZiB0aGUgaG9zdGVkIENob3JpYSBBdXRvbm9tb3VzIEFnZW50cyIsCiAgICAgICJhZ2dyZWdhdGUiOiBbCiAgICAgICAgewogICAgICAgICAgImZ1bmN0aW9uIjogInN1bW1hcnkiLAogICAgICAgICAgImFyZ3MiOiBbCiAgICAgICAgICAgICJtYWNoaW5lX25hbWVzIgogICAgICAgICAgXQogICAgICAgIH0KICAgICAgXQogICAgfSwKICAgIHsKICAgICAgImFjdGlvbiI6ICJtYWNoaW5lX3RyYW5zaXRpb24iLAogICAgICAiaW5wdXQiOiB7CiAgICAgICAgImluc3RhbmNlIjogewogICAgICAgICAgInByb21wdCI6ICJJbnN0YW5jZSBJRCIsCiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiTWFjaGluZSBJbnN0YW5jZSBJRCIsCiAgICAgICAgICAidHlwZSI6ICJzdHJpbmciLAogICAgICAgICAgImRlZmF1bHQiOiBudWxsLAogICAgICAgICAgIm9wdGlvbmFsIjogdHJ1ZSwKICAgICAgICAgICJ2YWxpZGF0aW9uIjogIl4uKy0uKy0uKy0uKy0uKyQiLAogICAgICAgICAgIm1heGxlbmd0aCI6IDM2CiAgICAgICAgfSwKICAgICAgICAidmVyc2lvbiI6IHsKICAgICAgICAgICJwcm9tcHQiOiAiVmVyc2lvbiIsCiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiTWFjaGluZSBWZXJzaW9uIiwKICAgICAgICAgICJ0eXBlIjogInN0cmluZyIsCiAgICAgICAgICAiZGVmYXVsdCI6IG51bGwsCiAgICAgICAgICAib3B0aW9uYWwiOiB0cnVlLAogICAgICAgICAgInZhbGlkYXRpb24iOiAiXlxcZCtcXC5cXGQrXFwuXFxkKyQiLAogICAgICAgICAgIm1heGxlbmd0aCI6IDIwCiAgICAgICAgfSwKICAgICAgICAibmFtZSI6IHsKICAgICAgICAgICJwcm9tcHQiOiAiTmFtZSIsCiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiTWFjaGluZSBOYW1lIiwKICAgICAgICAgICJ0eXBlIjogInN0cmluZyIsCiAgICAgICAgICAiZGVmYXVsdCI6IG51bGwsCiAgICAgICAgICAib3B0aW9uYWwiOiB0cnVlLAogICAgICAgICAgInZhbGlkYXRpb24iOiAiXlthLXpBLVpdW2EtekEtWjAtOV8tXSsiLAogICAgICAgICAgIm1heGxlbmd0aCI6IDEyOAogICAgICAgIH0sCiAgICAgICAgInBhdGgiOiB7CiAgICAgICAgICAicHJvbXB0IjogIlBhdGgiLAogICAgICAgICAgImRlc2NyaXB0aW9uIjogIk1hY2hpbmUgUGF0aCIsCiAgICAgICAgICAidHlwZSI6ICJzdHJpbmciLAogICAgICAgICAgImRlZmF1bHQiOiBudWxsLAogICAgICAgICAgIm9wdGlvbmFsIjogdHJ1ZSwKICAgICAgICAgICJ2YWxpZGF0aW9uIjogIi4rIiwKICAgICAgICAgICJtYXhsZW5ndGgiOiA1MTIKICAgICAgICB9LAogICAgICAgICJ0cmFuc2l0aW9uIjogewogICAgICAgICAgInByb21wdCI6ICJUcmFuc2l0aW9uIE5hbWUiLAogICAgICAgICAgImRlc2NyaXB0aW9uIjogIlRoZSB0cmFuc2l0aW9uIGV2ZW50IHRvIHNlbmQgdG8gdGhlIG1hY2hpbmUiLAogICAgICAgICAgInR5cGUiOiAic3RyaW5nIiwKICAgICAgICAgICJkZWZhdWx0IjogbnVsbCwKICAgICAgICAgICJvcHRpb25hbCI6IGZhbHNlLAogICAgICAgICAgInZhbGlkYXRpb24iOiAiXlthLXpBLVpdW2EtekEtWjAtOV8tXSskIiwKICAgICAgICAgICJtYXhsZW5ndGgiOiAxMjgKICAgICAgICB9CiAgICAgIH0sCiAgICAgICJvdXRwdXQiOiB7CiAgICAgICAgInN1Y2Nlc3MiOiB7CiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiSW5kaWNhdGVzIGlmIHRoZSB0cmFuc2l0aW9uIHdhcyBzdWNjZXNzZnVsbHkgYWNjZXB0ZWQiLAogICAgICAgICAgImRpc3BsYXlfYXMiOiAiQWNjZXB0ZWQiLAogICAgICAgICAgImRlZmF1bHQiOiBudWxsCiAgICAgICAgfQogICAgICB9LAogICAgICAiZGlzcGxheSI6ICJmYWlsZWQiLAogICAgICAiZGVzY3JpcHRpb24iOiAiQXR0ZW1wdHMgdG8gZm9yY2UgYSB0cmFuc2l0aW9uIGluIGEgaG9zdGVkIENob3JpYSBBdXRvbm9tb3VzIEFnZW50IgogICAgfQogIF0KfQo="

// DDLBytes is the raw JSON encoded DDL file for the agent
func DDLBytes() ([]byte, error) {
	return base64.StdEncoding.DecodeString(rawDDL)
}

// DDL is a parsed and loaded DDL for the agent
func DDL() (*agent.DDL, error) {
	ddlj, err := DDLBytes()
	if err != nil {
		return nil, err
	}

	ddl := &agent.DDL{}
	err = json.Unmarshal(ddlj, ddl)
	if err != nil {
		return nil, err
	}

	return ddl, nil
}