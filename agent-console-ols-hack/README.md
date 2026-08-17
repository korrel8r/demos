# Overview

This demo shows the troubleshooting panel working with OpenShift Lightspeed in the console.

The TP panel has been hacked to move it out of the way of the OLS panel,
the final release will have these panels co-exist properly.

# Setup

Using scenario: https://github.com/rhobs/troubleshooting-scenarios/blob/main/generic/01-payments-api-failure/

Setup: See generic/01-payments-api-failure/REAME.md: `make deploy; [wait 5m]; make break`

# Script

1. Start on home page, open lightspeed, select troubleshooting mode
2. Open TP, enable the console integration, Close TP
3. Open alerts (project payments, non-platform), open critical PaymentErrorRateHigh alert
4. Open OLS
5. "what am I looking at" - korrel8r provides better console communication.
6. "show me related information" (OLS both explains and displays a graph)
7. Click to logs, network events
8. Analyse the logs I'm looking at - root cause.
9. Conclusion: 
- seamless transition between chat and console navigation.
- lets user work together with chat to diagnose problems.


Other demo: compare
- find the root cause of this alert using correlation
- find the root cause of this alert, do not use any korrel8r tools

# NOTES

Korrel8r gives a better interpretation of the console URL than lightspeed currently has.
Lookin at: https://console-openshift-console.apps.snoflake.lan/monitoring/alertrules/718094864

Without korrel8r:

> what am I looking at
> OpenShift LightspeedAI
>
> Let me check what you're viewing in the console.I couldn't connect to your console directly, but based on the context, you're viewing AlertingRule "718094864". Let me look up the active alerts to give you more details.You're looking at an AlertingRule in the OpenShift web console. While I can't resolve the specific rule ID "718094864" to a named rule directly, 

> here's a summary of what's currently firing on your cluster. The AlertingRule page shows the definition of a Prometheus alerting rule — its PromQL expression, severity, labels, and which alerts it has produced.
>
> Your cluster currently has several active alerts across these categories

... proceeds to list all active alerts, not very helpful


With korrel8r:

> what am I looking at
> OpenShift LightspeedAI
>
> Let me check what you're viewing in the console and get the current alerts.You're looking at the PaymentErrorRateHigh alerting rule. Here's what it's telling you:
> 🔴 PaymentErrorRateHigh — Critical
>
>     Namespace: payments
>     Status: Actively firing
>     Error rate: 100% (threshold: 15%)
>     Firing since: Aug 10, 19:35 UTC
>     Description: The payment service has a 100% error rate — every payment request is failing.
>
> There's also a warning-level instance of this same rule (threshold: 3%) that is currently suppressed/inhibited by the critical alert, which is expected behavior.
> What this means
>
> The payment service in the payments namespace is completely failing. No successful payment transactions are going through.
>
> Would you like me to investigate the root cause? I can:
>
>     Check the pods and logs in the payments namespace
>     Look at the underlying metrics driving this alert
>     Check for related issues (e.g., the PostgreSQL connection problems in shared-services could be a contributing factor)

