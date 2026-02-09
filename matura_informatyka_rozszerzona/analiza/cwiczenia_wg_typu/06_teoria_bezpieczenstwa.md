# 06. Teoria bezpieczenstwa

Typ zadania: **teoria_bezpieczenstwa**
Czestotliwosc: 2/11 lat | Laczna punktacja: 2 pkt
Kategoria: TEORIA

---

### Cwiczenie 6.1 (trudnosc: latwe, ~1 pkt)
**Zrodlo inspiracji**: Matura 2025 zad. 4 (Keylogger)

Dopasuj opis sytuacji do odpowiedniego typu zagrozenia. Wpisz w tabeli nazwe zagrozenia wybrana ze zbioru: {keylogger, phishing, ransomware, trojan}.

| Lp. | Opis sytuacji | Typ zagrozenia |
|-----|---------------|----------------|
| a) | Uzytkownik otrzymal e-mail wyglodajacy jak wiadomosc z banku z prosba o klikniecie linku i podanie loginu oraz hasla. Strona pod linkiem wyglada identycznie jak strona banku, ale ma inny adres URL. | |
| b) | Po zainstalowaniu darmowego programu do edycji zdjec komputer zaczal wysylac dane uzytkownika na nieznany serwer. Program wyglodal na legalne oprogramowanie. | |
| c) | Wszystkie pliki na dysku uzytkownika zostaly zaszyfrowane. Na ekranie pojawil sie komunikat z zodaniem zaplacenia 500 dolarow w kryptowalucie za odzyskanie dostpu do danych. | |
| d) | Na komputerze znaleziono program, ktory rejestruje kazde nacisniecie klawisza i wysyla te dane do zewnetrznego serwera. Program dzialal w tle, niewidoczny dla uzytkownika. | |

<details>
<summary>Odpowiedz</summary>

| Lp. | Typ zagrozenia | Uzasadnienie |
|-----|----------------|-------------|
| a) | **phishing** | Phishing to metoda wyludzania danych (loginow, hasel, numerow kart) poprzez podszywanie sie pod zaufana instytucje. Charakterystyczny element: falszywa strona imitujaca prawdziwa. |
| b) | **trojan** (kon trojanski) | Trojan to zlosliwe oprogramowanie ukryte wewnatrz pozornie przydatnego programu. Uzytkownik sam go instaluje, nie wiedzac o ukrytych funkcjach (np. kradzez danych). |
| c) | **ransomware** | Ransomware szyfruje pliki ofiary i zoda okupu (ransom) za klucz deszyfrujacy. Czesto zodanie jest w kryptowalucie, by utrudnic sledzenie platnosci. |
| d) | **keylogger** | Keylogger to program rejestrujacy nacisniecia klawiszy (key = klawisz, logger = rejestrator). Pozwala przechwycic hasla, wiadomosci i inne wrdzliwe dane wpisywane z klawiatury. |
</details>

---

### Cwiczenie 6.2 (trudnosc: latwe, ~1 pkt)
**Zrodlo inspiracji**: Matura 2024 zad. 5 (Protokoly)

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

---

### Cwiczenie 6.3 (trudnosc: srednie, ~2 pkt)
**Zrodlo inspiracji**: Matura 2023 zad. 4 (szyfrowanie asymetryczne)

Ocen prawdziwosc ponizszych zdan dotyczacych kryptografii. Wpisz P (prawda) lub F (falsz).

| Lp. | Zdanie | P/F |
|-----|--------|-----|
| a) | W szyfrowaniu asymetrycznym do szyfrowania i deszyfrowania uzywa sie tego samego klucza. | |
| b) | Podpis cyfrowy jest tworzony z uzyciem klucza prywatnego nadawcy. | |
| c) | AES (Advanced Encryption Standard) jest algorytmem szyfrowania asymetrycznego. | |
| d) | Protokol HTTPS wykorzystuje poloczenie szyfrowania symetrycznego (do przesylania danych) i asymetrycznego (do uzgodnienia klucza sesji). | |

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

---

### Cwiczenie 6.4 (trudnosc: srednie, ~2 pkt)
**Zrodlo inspiracji**: styl mieszany 2023-2025

Ocen prawdziwosc ponizszych zdan dotyczacych bezpieczenstwa sieci i danych. Wpisz P (prawda) lub F (falsz).

| Lp. | Zdanie | P/F |
|-----|--------|-----|
| a) | Zapora sieciowa (firewall) blokuje caly ruch sieciowy przychodzacy do komputera. | |
| b) | VPN (Virtual Private Network) szyfruje transmisje danych miedzy urzodzeniem uzytkownika a serwerem VPN. | |
| c) | Certyfikat SSL/TLS potwierdza tozsamosc wlasciciela strony internetowej. | |
| d) | Haslo o dlugosci 8 znakow, zlozone wylacznie z cyfr, ma dokladnie 10^8 (sto milionow) mozliwych kombinacji. | |

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

---

### Cwiczenie 6.5 (trudnosc: trudne, ~3 pkt)
**Zrodlo inspiracji**: styl rozszerzony, analiza scenariusza

Przeczytaj opis scenariusza i odpowiedz na pytania.

**Scenariusz**: Pracownik firmy otrzymal e-mail od nadawcy podajacego sie za dzial IT firmy. W wiadomosci informowano o "krytycznej aktualizacji systemu bezpieczenstwa", ktora nalezy zainstalowac natychmiast. E-mail zawieral link do pobrania pliku "aktualizacja_v2.exe". Pracownik kliknal link, pobral i uruchomil plik. Po kilku dniach okazalo sie, ze z serwera firmowego wykradziono baze danych klientow.

**Polecenie**:
- a) Jaki typ ataku socjotechnicznego zostal zastosowany? Uzasadnij odpowiedz.
- b) Jaki rodzaj zlosliwego oprogramowania mogl zostac zainstalowany? Podaj nazwe i krotkie uzasadnienie.
- c) Wymien 3 srodki ochrony, ktore moglyby zapobiec temu incydentowi.

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
