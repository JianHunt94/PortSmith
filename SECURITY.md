# Security Policy

## Reporting a vulnerability

Please do not open a public issue for a security problem. Use GitHub's private
[security advisory](https://github.com/haebibisimax/PortSmith/security/advisories/new)
form, or contact a maintainer directly if the form is unavailable to you.

A useful report includes:

- what an attacker gains, in one sentence,
- the version or commit you are running,
- the smallest reproduction you can manage, and
- whether the input is attacker-controlled in your setup.

You should get an acknowledgement within a few days and a disclosure timeline
before anything is published.

## Scope

PortSmith reads recorded traffic and infers collectors from it. The capture files
and the intermediate trace format are both untrusted input: a truncated record, a
length field that disagrees with the payload, or a session id that is not validated
are the shapes worth probing.
