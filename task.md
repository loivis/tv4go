# Arbetsprov

Arbetsuppgiften är främst tänkt som ett underlag för att diskutera kod och lösningar. Du kommer få dela med dig av dina tankar kring koden vid en mobbprogrammeringssituation och vidareutveckla lösningen ihop med teamet.

Använd gärna Go och ägna max några timmar åt denna uppgift.

## Bedömning

* Fokusera på kodens struktur och kvalitet då detta är underlag för sessionen med teamet.

## Uppgift

I tjänsten C-More tillåter vi idag upp till två simultana videoströmmar, d.v.s. en användare kan se på två videor samtidigt.

* Bygg en mikrotjänst som en klient (web, mobil, ...) kan använda för att rapportera att en användare startat eller stoppat en videoström. Exempelvis genom request till endpoints `/startstream` resp. `/stopstream`. Spara denna information på något sätt.

* Validera att den video som rapporteras är giltig genom att söka efter videon i C-Mores sök-tjänst https://cmore-search.b17g.services/. Ett video-ID som inte existerar är ogiltigt och ska resultera i att ett fel returneras till klienten.

## Krav

* Det ska finnas minst ett enhetstest

* Det ska finnas API-dokumentation

* Tester och kod ska vara körbara
