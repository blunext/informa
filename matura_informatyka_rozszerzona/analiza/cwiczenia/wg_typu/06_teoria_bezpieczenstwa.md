# 06. Teoria bezpieczenstwa

Typ zadania: **teoria_bezpieczenstwa**
Czestotliwosc: 2/12 lat | Laczna punktacja: 2 pkt
Kategoria: TEORIA

## Umiejetnosci cwiczone w tym zestawie

`phishing` `ransomware` `trojan` `keylogger` `protokoly-sieciowe` `HTTPS` `FTP` `DNS` `DHCP` `SMTP` `szyfrowanie-symetryczne` `szyfrowanie-asymetryczne` `podpis-cyfrowy` `firewall` `VPN` `certyfikat-SSL` `socjotechnika` `atak-brute-force` `2FA` `RODO`

---

### Cwiczenie 6.1 (trudnosc: latwe, ~1 pkt)
**Zrodlo inspiracji**: Matura 2025 zad. 4 (Keylogger)
**Tagi**: `phishing` `trojan` `ransomware` `keylogger`

Dopasuj opis sytuacji do odpowiedniego typu zagrozenia. Wpisz w tabeli nazwe zagrozenia wybrana ze zbioru: {keylogger, phishing, ransomware, trojan}.

| Lp. | Opis sytuacji | Typ zagrozenia |
|-----|---------------|----------------|
| a) | Uzytkownik otrzymal e-mail wyglodajacy jak wiadomosc z banku z prosba o klikniecie linku i podanie loginu oraz hasla. Strona pod linkiem wyglada identycznie jak strona banku, ale ma inny adres URL. | |
| b) | Po zainstalowaniu darmowego programu do edycji zdjec komputer zaczal wysylac dane uzytkownika na nieznany serwer. Program wyglodal na legalne oprogramowanie. | |
| c) | Wszystkie pliki na dysku uzytkownika zostaly zaszyfrowane. Na ekranie pojawil sie komunikat z zodaniem zaplacenia 500 dolarow w kryptowalucie za odzyskanie dostpu do danych. | |
| d) | Na komputerze znaleziono program, ktory rejestruje kazde nacisniecie klawisza i wysyla te dane do zewnetrznego serwera. Program dzialal w tle, niewidoczny dla uzytkownika. | |

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Przypomnij sobie definicje: phishing = wyludzanie danych, trojan = zlosliwy program udajacy legalne oprogramowanie, ransomware = szyfrowanie z zadaniem okupu, keylogger = rejestracja klawiszy.
2. **Podejscie**: Zwroc uwage na slowa kluczowe: "e-mail z banku" = phishing, "darmowy program" = trojan, "zaszyfrowane pliki" = ransomware, "rejestruje nacisniecia klawiszy" = keylogger.
3. **Kluczowy krok**: Phishing to technika socjotechniczna (nie oprogramowanie). Trojan, ransomware i keylogger to rodzaje malware.

</details>

<details>
<summary>Odpowiedz</summary>

| Lp. | Typ zagrozenia | Uzasadnienie |
|-----|----------------|-------------|
| a) | **phishing** | Phishing to metoda wyludzania danych (loginow, hasel, numerow kart) poprzez podszywanie sie pod zaufana instytucje. Charakterystyczny element: falszywa strona imitujaca prawdziwa. |
| b) | **trojan** (kon trojanski) | Trojan to zlosliwe oprogramowanie ukryte wewnatrz pozornie przydatnego programu. Uzytkownik sam go instaluje, nie wiedzac o ukrytych funkcjach (np. kradzez danych). |
| c) | **ransomware** | Ransomware szyfruje pliki ofiary i zoda okupu (ransom) za klucz deszyfrujacy. Czesto zodanie jest w kryptowalucie, by utrudnic sledzenie platnosci. |
| d) | **keylogger** | Keylogger to program rejestrujacy nacisniecia klawiszy (key = klawisz, logger = rejestrator). Pozwala przechwycic hasla, wiadomosci i inne wrazliwe dane wpisywane z klawiatury. |
</details>

<details>
<summary>Typowe bledy</summary>

- **Trojan = wirus**: Trojan NIE jest wirusem (wirus sie replikuje, trojan nie). CKE: -0.5 pkt
- **Phishing = malware**: Phishing to technika socjotechniczna, nie oprogramowanie. CKE: -0.5 pkt
- **Ransomware = keylogger**: Ransomware szyfruje pliki, keylogger rejestruje klawisze — zupelnie rozne dzialanie. CKE: -0.5 pkt

</details>

---

### Cwiczenie 6.2 (trudnosc: latwe, ~1 pkt)
**Zrodlo inspiracji**: Matura 2024 zad. 5 (Protokoly)
**Tagi**: `protokoly-sieciowe` `HTTPS` `FTP` `DNS` `DHCP` `SMTP`

Dopasuj protokol sieciowy do jego funkcji. Wpisz w tabeli nazwe protokolu wybrana ze zbioru: {HTTP, FTP, DHCP, DNS, SMTP, HTTPS}.

| Lp. | Opis funkcji | Protokol |
|-----|-------------|----------|
| a) | Automatyczne przydzielanie adresow IP urzodzeniom w sieci lokalnej. | |
| b) | Przesylanie plikow miedzy komputerem a serwerem (upload i download plikow). | |
| c) | Zamienianie nazw domenowych (np. www.google.com) na adresy IP. | |
| d) | Przesylanie stron internetowych z szyfrowaniem danych (poloczenie bezpieczne). | |
| e) | Przesylanie stron internetowych bez szyfrowania. | |
| f) | Wysylanie poczty elektronicznej z klienta do serwera pocztowego. | |

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Kazdy protokol ma jedna glowna funkcje. Skojarz akronimy: D=Dynamic, H=Host, C=Configuration, P=Protocol.
2. **Podejscie**: HTTP vs HTTPS — jedyna roznica to szyfrowanie (S = Secure). FTP = pliki, SMTP = poczta.
3. **Kluczowy krok**: DNS zamienia domeny na IP. DHCP przydziela adresy IP automatycznie.

</details>

<details>
<summary>Odpowiedz</summary>

| Lp. | Protokol | Uzasadnienie |
|-----|----------|-------------|
| a) | **DHCP** | Dynamic Host Configuration Protocol — automatycznie przydziela adresy IP, maski podsieci i inne parametry konfiguracyjne urzodzeniom w sieci. |
| b) | **FTP** | File Transfer Protocol — protokol przesylania plikow. Umozliwia upload (wysylanie) i download (pobieranie) plikow z/na serwer. |
| c) | **DNS** | Domain Name System — tlumaczenie nazw domenowych na adresy IP (np. www.google.com -> 142.250.185.14). |
| d) | **HTTPS** | HyperText Transfer Protocol Secure — wersja HTTP z szyfrowaniem TLS/SSL. Przesyla strony internetowe w sposob bezpieczny. |
| e) | **HTTP** | HyperText Transfer Protocol — podstawowy protokol przesylania stron www, bez szyfrowania. |
| f) | **SMTP** | Simple Mail Transfer Protocol — protokol wysylania poczty elektronicznej od klienta do serwera i miedzy serwerami. |
</details>

<details>
<summary>Typowe bledy</summary>

- **DNS = DHCP**: DNS zamienia domeny na IP, DHCP przydziela adresy IP. Rozne funkcje! CKE: -0.5 pkt
- **HTTP = FTP**: HTTP przesyla strony WWW, FTP przesyla pliki. CKE: -0.5 pkt
- **SMTP = POP3/IMAP**: SMTP wysyla poczta, POP3/IMAP odbieraja. CKE: -0.5 pkt

</details>

---

### Cwiczenie 6.3 (trudnosc: srednie, ~2 pkt)
**Zrodlo inspiracji**: Matura 2023 zad. 4 (szyfrowanie asymetryczne)
**Tagi**: `szyfrowanie-symetryczne` `szyfrowanie-asymetryczne` `podpis-cyfrowy` `HTTPS`

Ocen prawdziwosc ponizszych zdan dotyczacych kryptografii. Wpisz P (prawda) lub F (falsz).

| Lp. | Zdanie | P/F |
|-----|--------|-----|
| a) | W szyfrowaniu asymetrycznym do szyfrowania i deszyfrowania uzywa sie tego samego klucza. | |
| b) | Podpis cyfrowy jest tworzony z uzyciem klucza prywatnego nadawcy. | |
| c) | AES (Advanced Encryption Standard) jest algorytmem szyfrowania asymetrycznego. | |
| d) | Protokol HTTPS wykorzystuje poloczenie szyfrowania symetrycznego (do przesylania danych) i asymetrycznego (do uzgodnienia klucza sesji). | |

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Asymetryczne = dwa klucze (publiczny + prywatny). Symetryczne = jeden klucz.
2. **Podejscie**: Podpis cyfrowy: tworzymy kluczem prywatnym, weryfikujemy publicznym. AES = symetryczny.
3. **Kluczowy krok**: HTTPS = kryptografia hybrydowa: asymetryczna do uzgodnienia klucza, symetryczna do danych.

</details>

<details>
<summary>Odpowiedz</summary>

| Lp. | Odpowiedz | Uzasadnienie |
|-----|-----------|-------------|
| a) | **F** | W szyfrowaniu asymetrycznym uzywa sie pary kluczy: klucza publicznego (do szyfrowania) i klucza prywatnego (do deszyfrowania). To wlasnie szyfrowanie SYMETRYCZNE uzywa jednego wspolnego klucza do obu operacji. |
| b) | **P** | Podpis cyfrowy jest tworzony przez nadawce za pomoca jego klucza prywatnego. Odbiorca weryfikuje podpis uzywajac klucza publicznego nadawcy. Gwarantuje to autentycznosc i niezmiennosc wiadomosci. |
| c) | **F** | AES jest algorytmem szyfrowania SYMETRYCZNEGO (ten sam klucz do szyfrowania i deszyfrowania). Przykladami algorytmow asymetrycznych sa RSA i ECC. |
| d) | **P** | HTTPS (TLS) stosuje kryptografie hybrydowa: najpierw uzgadnia klucz sesji za pomoca algorytmu asymetrycznego (np. RSA, ECDH), a nastepnie szyfruje wlosciwa komunikacje szybszym algorytmem symetrycznym (np. AES). |

Odpowiedzi: a) F, b) P, c) F, d) P
</details>

<details>
<summary>Typowe bledy</summary>

- **Asymetryczne = jeden klucz**: Asymetryczne = DWA klucze. Symetryczne = JEDEN klucz. CKE: -0.5 pkt
- **AES = asymetryczny**: AES jest SYMETRYCZNY. RSA i ECC sa asymetryczne. CKE: -0.5 pkt
- **Podpis cyfrowy = klucz publiczny**: Podpis TWORZYMY kluczem prywatnym, WERYFIKUJEMY publicznym. CKE: -0.5 pkt

</details>

---

### Cwiczenie 6.4 (trudnosc: srednie, ~2 pkt)
**Zrodlo inspiracji**: styl mieszany 2023-2025
**Tagi**: `firewall` `VPN` `certyfikat-SSL` `atak-brute-force`

Ocen prawdziwosc ponizszych zdan dotyczacych bezpieczenstwa sieci i danych. Wpisz P (prawda) lub F (falsz).

| Lp. | Zdanie | P/F |
|-----|--------|-----|
| a) | Zapora sieciowa (firewall) blokuje caly ruch sieciowy przychodzacy do komputera. | |
| b) | VPN (Virtual Private Network) szyfruje transmisje danych miedzy urzodzeniem uzytkownika a serwerem VPN. | |
| c) | Certyfikat SSL/TLS potwierdza tozsamosc wlasciciela strony internetowej. | |
| d) | Haslo o dlugosci 8 znakow, zlozone wylacznie z cyfr, ma dokladnie 10^8 (sto milionow) mozliwych kombinacji. | |

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Firewall filtruje ruch (nie blokuje calego). VPN tworzy szyfrowany tunel.
2. **Podejscie**: Certyfikat SSL potwierdza, ze domena nalezy do podmiotu. 10 cyfr ^ 8 pozycji.
3. **Kluczowy krok**: 10^8 = 100 mln. To bardzo malo — takie haslo mozna zlamac brute-force w sekundy.

</details>

<details>
<summary>Odpowiedz</summary>

| Lp. | Odpowiedz | Uzasadnienie |
|-----|-----------|-------------|
| a) | **F** | Firewall NIE blokuje calego ruchu przychodzacego — filtruje ruch na podstawie zdefiniowanych regul. Przepuszcza ruch dozwolony (np. odpowiedzi na Twoje zapytania HTTP) i blokuje ruch podejrzany lub nieautoryzowany. |
| b) | **P** | VPN tworzy zaszyfrowany tunel miedzy urzodzeniem uzytkownika a serwerem VPN. Caly ruch sieciowy przechodzacy przez ten tunel jest szyfrowany, co chroni dane przed podsluchem (np. w publicznych sieciach Wi-Fi). |
| c) | **P** | Certyfikat SSL/TLS jest wydawany przez zaufany urzod certyfikacji (CA) i potwierdza, ze domena nalezy do okreslonego podmiotu. Przegladarka weryfikuje certyfikat i wyswietla ikone klocki przy adresie HTTPS. |
| d) | **P** | Kazdy z 8 znakow moze byc jedna z 10 cyfr (0-9). Liczba kombinacji: 10 * 10 * 10 * 10 * 10 * 10 * 10 * 10 = 10^8 = 100 000 000. Dlatego hasla zlozoone z samych cyfr sa slabe — mozna je zlamac atakiem brute-force. |

Odpowiedzi: a) F, b) P, c) P, d) P
</details>

<details>
<summary>Typowe bledy</summary>

- **Firewall = blokuje wszystko**: Firewall FILTRUJE ruch wg regul, nie blokuje calego. CKE: -0.5 pkt
- **VPN = calkowita anonimnosc**: VPN szyfruje polaczenie do serwera VPN, ale operator VPN widzi ruch. CKE: -0.5 pkt
- **10^8 = duzo**: 100 mln kombinacji to bardzo malo. Brute-force to trwa ulamek sekundy na nowoczesnym komputerze. CKE: brak kary, ale warto wiedziec.

</details>

---

### Cwiczenie 6.5 (trudnosc: trudne, ~3 pkt)
**Zrodlo inspiracji**: styl rozszerzony, analiza scenariusza
**Tagi**: `phishing` `trojan` `socjotechnika`

Przeczytaj opis scenariusza i odpowiedz na pytania.

**Scenariusz**: Pracownik firmy otrzymal e-mail od nadawcy podajacego sie za dzial IT firmy. W wiadomosci informowano o "krytycznej aktualizacji systemu bezpieczenstwa", ktora nalezy zainstalowac natychmiast. E-mail zawieral link do pobrania pliku "aktualizacja_v2.exe". Pracownik kliknal link, pobral i uruchomil plik. Po kilku dniach okazalo sie, ze z serwera firmowego wykradziono baze danych klientow.

**Polecenie**:
- a) Jaki typ ataku socjotechnicznego zostal zastosowany? Uzasadnij odpowiedz.
- b) Jaki rodzaj zlosliwego oprogramowania mogl zostac zainstalowany? Podaj nazwe i krotkie uzasadnienie.
- c) Wymien 3 srodki ochrony, ktore moglyby zapobiec temu incydentowi.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Atak polega na podszywaniu sie pod zaufane zrodlo (dzial IT) — to phishing. Program .exe udajacy aktualizacje = trojan.
2. **Podejscie**: Srodki ochrony: szkolenia + filtrowanie poczty + ograniczenie uprawnien.
3. **Kluczowy krok**: Spear phishing = ukierunkowany phishing (na konkretna firme/osobe, nie masowy).

</details>

<details>
<summary>Odpowiedz</summary>

**a) Typ ataku: phishing (spear phishing)**

Uzasadnienie: Atak polegal na podszywaniu sie pod zaufany podmiot (dzial IT firmy) w celu naklonenia ofiary do wykonania niebezpiecznej czynnosci (pobrania i uruchomienia zlosliwego pliku). Jest to odmiana phishingu zwana spear phishing, poniewaz wiadomosc byla ukierunkowana na konkretna osobe/firme (nie masowa wysylka). Elementy typowe dla phishingu:
- Podszywanie sie pod zaufane zrodlo (dzial IT)
- Wywieranie presji czasowej ("krytyczna aktualizacja", "natychmiast")
- Naklanianie do klikniecia linku i pobrania pliku

**b) Rodzaj zlosliwego oprogramowania: trojan (kon trojanski)**

Uzasadnienie: Plik "aktualizacja_v2.exe" wygladal jak legalne oprogramowanie (aktualizacja systemu), ale w rzeczywistosci zawieral zlosliwy kod. Trojan mogl zainstalowac backdoor (tylne drzwi) umozliwiajacy zdalny dostep do serwera firmowego, co pozwolilo na kradzez bazy danych klientow. Mozliwe, ze trojan zawieral rowniez keylogger do przechwycenia hasel dostepu do serwera.

**c) Trzy srodki ochrony:**

1. **Szkolenie pracownikow z rozpoznawania phishingu** — pracownik powinien wiedziec, ze dzial IT nie wysyla aktualizacji przez e-mail z linkami do pobrania. Powinien weryfikowac takie wiadomosci bezposrednio z dzialem IT (np. telefonicznie).

2. **Filtrowanie poczty i blokowanie zlosliwych zalacznikow** — system bezpieczenstwa poczty (anti-spam, anti-malware) powinien wykrywac podejrzane linki i pliki wykonywalne (.exe) w zalacznikach i blokowac je lub oznaczac jako niebezpieczne.

3. **Ograniczenie uprawnien uzytkownikow (zasada najmniejszych uprawnien)** — pracownik nie powinien miec mozliwosci instalowania oprogramowania bez zgody administratora. Oprogramowanie powinno byc instalowane wylacznie z autoryzowanych zrodel.

Dodatkowe srodki (alternatywne):
- Program antywirusowy z aktualna baza sygnatur
- Dwuskladnikowe uwierzytelnianie (2FA) do krytycznych systemow
- Segmentacja sieci — oddzielenie stacji roboczych od serwerow z danymi
</details>

<details>
<summary>Typowe bledy</summary>

- **Phishing zamiast spear phishing**: CKE akceptuje "phishing", ale "spear phishing" jest dokladniejszy (ukierunkowany). CKE: brak kary za ogolna nazwe.
- **Wirus zamiast trojan**: Trojan udaje legalne oprogramowanie. Wirus sie sam replikuje. CKE: -0.5 pkt
- **Tylko 2 srodki ochrony zamiast 3**: CKE wymaga dokladnie tyle ile wskazano w poleceniu. CKE: -1 pkt

</details>

---

### Cwiczenie 6.6 (trudnosc: latwe, ~1 pkt)
**Zrodlo inspiracji**: styl CKE, dopasowywanie pojec
**Tagi**: `socjotechnika` `phishing` `ransomware`

Dopasuj rodzaj ataku do opisu. Wybierz ze zbioru: {DDoS, man-in-the-middle, SQL injection, brute force}.

| Lp. | Opis ataku | Rodzaj |
|-----|-----------|--------|
| a) | Atakujacy wysyla ogromna liczbe zapytan do serwera, aby go przeciazyc i uniemozliwic dostep innym uzytkownikom. | |
| b) | Atakujacy przechwytuje komunikacje miedzy dwoma stronami (np. uzytkownikiem i bankiem) i moze ja odczytywac lub modyfikowac. | |
| c) | Atakujacy wstawia do formularza na stronie internetowej fragment kodu SQL, ktory pozwala mu odczytac dane z bazy danych serwera. | |
| d) | Atakujacy probuje wszystkie mozliwe kombinacje hasel po kolei, az znajdzie poprawne haslo. | |

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: DDoS = przeciazenie serwera. MITM = przechwycenie komunikacji. SQLi = atak na baze danych. Brute force = probowanie hasel.
2. **Podejscie**: Slowa kluczowe: "ogromna liczba zapytan" = DDoS, "przechwytuje" = MITM, "kod SQL" = SQL injection, "wszystkie kombinacje" = brute force.
3. **Kluczowy krok**: SQL injection to atak na APLIKACJE WEBOWE (formularze). DDoS to atak na DOSTEPNOSC (nie kradziezy danych).

</details>

<details>
<summary>Odpowiedz</summary>

| Lp. | Rodzaj ataku | Uzasadnienie |
|-----|-------------|-------------|
| a) | **DDoS** | Distributed Denial of Service — rozproszony atak polegajacy na zalaniu serwera ogromna liczba zapytan z wielu zrodel, co prowadzi do odmowy uslugi (serwer nie moze obslugiwac prawdziwych uzytkownikow). |
| b) | **man-in-the-middle** | Atak "czlowiek posrodku" — atakujacy umieszcza sie miedzy komunikujacymi sie stronami, przechwytujac i potencjalnie modyfikujac przesylane dane. Chroni przed tym szyfrowanie HTTPS. |
| c) | **SQL injection** | Wstrzykniecie SQL — atakujacy umieszcza zlosliwy kod SQL w polu formularza (np. login), ktory jest wykonywany przez baze danych serwera, umozliwiajac odczyt lub modyfikacje danych. |
| d) | **brute force** | Atak silowy — systematyczne probowanie wszystkich mozliwych kombinacji hasel. Dla hasla 4-cyfrowego to 10000 prob. Ochrona: blokada konta po kilku blednych probach, dluzsze hasla. |
</details>

<details>
<summary>Typowe bledy</summary>

- **DDoS = kradzez danych**: DDoS atakuje DOSTEPNOSC, nie poufnosc danych. CKE: -0.5 pkt
- **SQL injection = atak na siec**: SQLi to atak na APLIKACJE (warstwe oprogramowania), nie na siec. CKE: -0.5 pkt
- **MITM = phishing**: MITM przechwytuje istniejaca komunikacje. Phishing wyludza dane podszywajac sie. CKE: -0.5 pkt

</details>

---

### Cwiczenie 6.7 (trudnosc: srednie, ~2 pkt)
**Zrodlo inspiracji**: Matura 2023 zad. 4.2 (scenariusz bezpieczenstwa)
**Tagi**: `2FA` `szyfrowanie-asymetryczne` `certyfikat-SSL` `VPN`

Wyjasnij krotko (1-2 zdania) nastepujace pojecia z zakresu bezpieczenstwa informatycznego:

| Lp. | Pojecie |
|-----|---------|
| a)  | Uwierzytelnianie dwuskladnikowe (2FA) |
| b)  | Podpis cyfrowy |
| c)  | Certyfikat SSL/TLS |
| d)  | Atak typu zero-day |

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: 2FA = dwa niezalezne sposoby potwierdzenia tozsamosci. Podpis cyfrowy = gwarancja autentycznosci.
2. **Podejscie**: Certyfikat SSL = "dowod osobisty" strony internetowej. Zero-day = luka, o ktorej producent nie wie.
3. **Kluczowy krok**: 2FA: cos wiesz (haslo) + cos masz (telefon) LUB cos jestes (biometria).

</details>

<details>
<summary>Odpowiedz</summary>

**a) Uwierzytelnianie dwuskladnikowe (2FA):**
Metoda potwierdzania tozsamosci wymagajaca dwoch niezaleznych skladnikow: zazwyczaj "cos wiesz" (haslo) i "cos masz" (kod SMS, aplikacja authenticator, klucz U2F). Nawet jesli atakujacy pozna haslo, nie zaloguje sie bez drugiego skladnika.

**b) Podpis cyfrowy:**
Mechanizm kryptograficzny pozwalajacy potwierdzic autentycznosc i integralnosc dokumentu elektronicznego. Nadawca tworzy podpis uzywajac swojego klucza prywatnego, a odbiorca weryfikuje go kluczem publicznym nadawcy — dzieki temu mozna potwierdzic, ze dokument pochodzi od konkretnej osoby i nie zostal zmieniony.

**c) Certyfikat SSL/TLS:**
Cyfrowy dokument wydawany przez zaufany urzad certyfikacji (CA), ktory potwierdza tozsamosc wlasciciela strony internetowej (ze domena nalezy do okreslonego podmiotu). Umozliwia nawiazanie zaszyfrowanego polaczenia HTTPS miedzy przegladarka a serwerem.

**d) Atak typu zero-day:**
Atak wykorzystujacy luke w oprogramowaniu, o ktorej producent (tworca oprogramowania) jeszcze nie wie lub na ktora nie wydal jeszcze poprawki (patch). Nazwa "zero-day" oznacza, ze producent mial zero dni na naprawienie bledu od momentu jego odkrycia przez atakujacego.
</details>

<details>
<summary>Typowe bledy</summary>

- **2FA = dwa hasla**: NIE! 2FA wymaga dwoch ROZNYCH typow skladnikow (cos wiesz + cos masz), nie dwoch hasel. CKE: -0.5 pkt
- **Podpis cyfrowy = szyfrowanie**: Podpis potwierdza AUTENTYCZNOSC, nie szyfruje tresci. CKE: -0.5 pkt
- **Zero-day = stara znana luka**: Zero-day to luka NIEZNANA producentowi. Znana luka to CVE. CKE: -0.5 pkt

</details>

---

### Cwiczenie 6.8 (trudnosc: srednie, ~2 pkt)
**Zrodlo inspiracji**: styl CKE, sila hasel
**Tagi**: `atak-brute-force` `szyfrowanie-symetryczne`

**Polecenie**: Oblicz i porownaj sile dwoch hasel:

- Haslo A: 6 znakow, tylko male litery (a-z), np. "abcdef"
- Haslo B: 4 znaki, male i duze litery + cyfry (a-z, A-Z, 0-9), np. "aB3x"

Podaj:
- a) Ile roznych hasel mozna utworzyc dla kazdego typu?
- b) Ktore haslo jest silniejsze (trudniejsze do zlamania brute-force)?
- c) Jezeli komputer probuje 10^9 (miliard) hasel na sekunde, ile czasu zajmie sprawdzenie wszystkich hasel typu A i typu B?

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Ilosc hasel = (liczba znakow)^(dlugosc). Male litery = 26, duze = 26, cyfry = 10.
2. **Podejscie**: A: 26^6. B: 62^4. Porownaj obie wartosci.
3. **Kluczowy krok**: 26^6 ≈ 309 mln. 62^4 ≈ 14.8 mln. Dluzsze haslo z mniejszym alfabetem moze byc silniejsze!

</details>

<details>
<summary>Odpowiedz</summary>

**a) Liczba roznych hasel:**

Haslo A: 26^6 = 26 * 26 * 26 * 26 * 26 * 26 = **308 915 776** (≈ 309 mln)

Haslo B: 62^4 = 62 * 62 * 62 * 62 = **14 776 336** (≈ 14.8 mln)

**b) Haslo A jest silniejsze** — ma ponad 20 razy wiecej mozliwych kombinacji.

Wniosek: dlugosc hasla ma WIEKSZY wplyw na sile niz roznorodnosc znakow. 6 malych liter > 4 znaki z duzego alfabetu.

**c) Czas lamania przy 10^9 hasel/s:**

Haslo A: 308 915 776 / 10^9 ≈ **0.31 sekundy**

Haslo B: 14 776 336 / 10^9 ≈ **0.015 sekundy** (15 milisekund)

Oba hasla sa BARDZO slabe przy wspolczesnej mocy obliczeniowej! Bezpieczne haslo powinno miec co najmniej 12 znakow z duzego alfabetu: 62^12 ≈ 3.2 * 10^21, co daje ~100 lat lamania.
</details>

<details>
<summary>Typowe bledy</summary>

- **26+26+10 = 62 zamiast 26*26*10**: Liczymy ILOCZYN (kazda pozycja niezalezna), nie sume. CKE: -1 pkt
- **Krotsze haslo silniejsze bo roznorodne**: Dlugosc wplywa wykladniczo, roznorodnosc liniowo. CKE: -0.5 pkt
- **Brak jednostki czasu**: Podaj wynik w sekundach/minutach/godzinach. CKE: -0.5 pkt

</details>

---

### Cwiczenie 6.9 (trudnosc: srednie-trudne, ~2 pkt)
**Zrodlo inspiracji**: Matura 2025 zad. 4.2, analiza polityki bezpieczenstwa
**Tagi**: `RODO` `firewall` `2FA` `socjotechnika`

Firma XYZ przechowuje dane osobowe klientow. Zaplanuj polityne bezpieczenstwa odpowiadajac na pytania:

- a) Wymien 3 organizacyjne srodki ochrony danych osobowych.
- b) Wymien 3 techniczne srodki ochrony danych osobowych.
- c) Czym jest zasada minimalnych uprawnien (principle of least privilege) i dlaczego jest wazna?

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Organizacyjne = szkolenia, polityki, procedury. Techniczne = szyfrowanie, firewall, backup.
2. **Podejscie**: RODO wymaga zarowno srodkow organizacyjnych, jak i technicznych.
3. **Kluczowy krok**: Minimalne uprawnienia = kazdy uzytkownik ma dostep TYLKO do tego, co potrzebuje do pracy.

</details>

<details>
<summary>Odpowiedz</summary>

**a) 3 organizacyjne srodki ochrony:**

1. **Szkolenia pracownikow** z zakresu bezpieczenstwa informacji i rozpoznawania atakow socjotechnicznych (phishing)
2. **Polityka hasel** — wymogi dotyczace dlugosci, zlozonosci i regularnej zmiany hasel
3. **Procedury reagowania na incydenty** — plan dzialania w razie wycieku danych (kogo powiadomic, jak zabezpieczyc dowody, jak zminimalizowac szkody)

**b) 3 techniczne srodki ochrony:**

1. **Szyfrowanie danych** — zarowno w spoczynku (na dyskach) jak i w tranzycie (HTTPS/TLS)
2. **Zapora sieciowa (firewall)** i system wykrywania wlaman (IDS/IPS)
3. **Regularne kopie zapasowe (backup)** — przechowywane w oddzielnej lokalizacji, testowane pod katem mozliwosci odtworzenia

**c) Zasada minimalnych uprawnien:**

Kazdy uzytkownik, program lub proces powinien miec dostep TYLKO do tych zasobow i danych, ktore sa niezbedne do wykonywania jego zadan — i nic wiecej.

Dlaczego jest wazna:
- Ogranicza szkody w razie przejecia konta (atakujacy uzyskuje dostep tylko do czesci systemu)
- Zmniejsza ryzyko przypadkowego uszkodzenia lub usuniecia danych
- Ulatwia audyt i sledzenie, kto mial dostep do czego
- Chroni przed atakami typu eskalacja uprawnien

Przyklad: Pracownik dzialu marketingu nie powinien miec dostepu do bazy danych finansowych firmy.
</details>

<details>
<summary>Typowe bledy</summary>

- **Tylko srodki techniczne bez organizacyjnych**: RODO wymaga OBU typow srodkow. CKE: -0.5 pkt
- **Brak przykladow**: CKE ceni konkretne przyklady, nie ogolniki. CKE: -0.5 pkt
- **Minimalne uprawnienia = brak dostepu**: NIE — chodzi o WYSTARCZAJACY dostep, nie zerowy. CKE: -0.5 pkt

</details>

---

### Cwiczenie 6.10 (trudnosc: trudne, ~3 pkt)
**Zrodlo inspiracji**: styl rozszerzony, analiza scenariusza z wyciekiem
**Tagi**: `socjotechnika` `szyfrowanie-asymetryczne` `phishing` `2FA`

**Scenariusz**: Bank "SafeBank" oferuje bankowosc internetowa. Klienci loguja sie haslem i kodem SMS (2FA). W zeszlym tygodniu:
1. Klient Marek otrzymal SMS: "Twoje konto zostalo zablokowane. Kliknij link aby je odblokowac: safebank-security.com". Link prowadzil do strony identycznie wygladajacej jak strona banku.
2. Marek wpisal swoj login, haslo i kod SMS na faltyszwej stronie.
3. Atakujacy natychmiast uzylyj tych danych do zalogowania sie na prawdziwej stronie banku i przelal pieniadze.

**Polecenie**:
- a) Jaki typ ataku zostal przeprowadzony? Dlaczego 2FA (kod SMS) nie ochronilo Marka?
- b) Podaj 2 slabosci 2FA opartego na kodach SMS.
- c) Zaproponuj bezpieczniejsza metode 2FA, ktora ochronilaby Marka, i wyjasnij dlaczego.
- d) Jak Marek mogl rozpoznac, ze strona jest faltywa?

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Atak to "real-time phishing" — atakujacy uzywa danych w czasie rzeczywistym, zanim kod SMS wygasnie.
2. **Podejscie**: Slabosci SMS 2FA: przechwycenie SIM (SIM swapping), phishing w czasie rzeczywistym, SMS nieszyfrowany.
3. **Kluczowy krok**: Klucze U2F/FIDO2 sa odporne na phishing, bo weryfikuja domene serwera automatycznie.

</details>

<details>
<summary>Odpowiedz</summary>

**a) Typ ataku: phishing w czasie rzeczywistym (real-time phishing / reverse proxy phishing)**

2FA oparte na kodach SMS NIE ochronilo Marka, poniewaz:
- Atakujacy uruchomil falszwa strone, ktora dzialala jako "proxy" — natychmiast przekazywal wprowadzone dane (login, haslo, kod SMS) na prawdziwa strone banku
- Kod SMS jest wazny przez 30-60 sekund — wystarczajaco dlugo, aby atakujacy go uzyj
- 2FA przez SMS chroni przed KRADZIEZIA HASLA (atakujacy potrzebuje tez kodu), ale NIE chroni przed phishingiem w czasie rzeczywistym

**b) 2 slabosci 2FA opartego na SMS:**

1. **Podatnosc na atak SIM swapping** — atakujacy moze przekonac operatora do przeniesienia numeru telefonu na swoja karte SIM, przejmujac wszystkie kody SMS
2. **Podatnosc na phishing w czasie rzeczywistym** — jak w scenariuszu, atakujacy moze natychmiast uzyc kodu SMS, zanim ten wygasnie. SMS nie weryfikuje, KOMU uzytkownik podaje kod

**c) Bezpieczniejsza metoda: klucz FIDO2/U2F (np. YubiKey)**

Klucze sprzetowe FIDO2 sa odporne na phishing, poniewaz:
- Klucz automatycznie weryfikuje DOMENE serwera (sprawdza, czy to naprawde safebank.pl, a nie safebank-security.com)
- Nawet jesli uzytkownik kliknie faltywy link, klucz NIE wygeneruje odpowiedzi dla blednej domeny
- Nie ma kodu do "przepisania" — komunikacja odbywa sie kryptograficznie miedzy kluczem a serwerem

Alternatywa: aplikacje TOTP (Google Authenticator) — nieco bezpieczniejsze niz SMS (nie podatne na SIM swap), ale nadal podatne na phishing w czasie rzeczywistym.

**d) Jak rozpoznac faltywa strone:**

1. **Sprawdzic adres URL** — safebank-security.com != safebank.pl. Prawdziwa domena banku to np. safebank.pl
2. **Sprawdzic certyfikat SSL** — kliknac ikone klocki w przegladarce i zweryfikowac, do kogo nalezy certyfikat
3. **Nie kliac linkow z SMS/email** — zawsze wchodzic na strone banku recznie (wpisac adres lub uzyc zakladki)
4. **Zwrocic uwage na jezyk** — wiadomosci phishingowe czesto zawieraja bledy jezykowe lub nietypowe sformulowania
</details>

<details>
<summary>Typowe bledy</summary>

- **2FA = pelne bezpieczenstwo**: 2FA przez SMS NIE chroni przed phishingiem real-time. CKE: -1 pkt
- **FIDO2 = to samo co SMS**: FIDO2 weryfikuje domene automatycznie, SMS nie. CKE: -0.5 pkt
- **Brak konkretnych slabosci SMS**: "SMS jest slaby" bez uzasadnienia to za malo. Podaj SIM swap i real-time phishing. CKE: -1 pkt

</details>

---

## Samoocena

Po rozwiazaniu cwiczen bez podgladania odpowiedzi, okresl swoj poziom:

| Poziom | Opis | Wynik |
|--------|------|-------|
| Podstawowy | Rozpoznajesz podstawowe zagrozenia (phishing, trojan, ransomware) i protokoly sieciowe | 1-3 cwiczen bez pomocy |
| Dobry | Rozumiesz roznice miedzy szyfrowaniem symetrycznym a asymetrycznym, znasz zasady 2FA | 4-6 cwiczen bez pomocy |
| Bardzo dobry | Analizujesz scenariusze bezpieczenstwa, obliczasz sile hasel, znasz srodki ochrony | 7-8 cwiczen bez pomocy |
| Doskonaly | Rozumiesz slabosci 2FA, FIDO2, real-time phishing i potrafisz zaplanowac polityke bezpieczenstwa | 9-10 cwiczen bez pomocy |

**Co dalej?**
- Poziom Podstawowy: Przerob cwiczenia 6.1, 6.2, 6.6 jeszcze raz. Wrocz do `cheatsheet_teoria.md` (sekcja: bezpieczenstwo).
- Poziom Dobry: Skup sie na cwiczeniach 6.3, 6.4, 6.7, 6.8. Przejdz do `04_test_prawda_falsz.md`.
- Poziom Bardzo dobry/Doskonaly: Przejdz do `05_konwersja_systemow_liczbowych.md` i `01_sledzenie_algorytmu.md`.
