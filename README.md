# Tester invio email

In questo repository è presente un piccolo tools che permette di testare/inviare email creando varie configurazioni.

## Build
Per creare la build del codice è necessario eseguire il seguente comando:

```bash
make build
```

Se la build ha successo, si otterrà un file chiamato `go-mail-local` nella cartella `bin`.

## Usage

### Creazione configurazione
Per creare una configurazione è necessario eseguire il seguente comando:

```bash
./bin/go-mail-local [-profile=] setup
```
Che crea una configurazione per inviare successivamente le email. Se non è presente il profilo
allora viene creato il profilo di default.

**N.B. Le configurazioni verranno salvate nella home directory dell'utente nella cartella .go-mail-local/config**

### Recupera configurazione
Per recuperare una configurazione è necessario eseguire il seguente comando:

```bash
./bin/go-mail-local [-profile=] setup get
```
Mostra la configurazione salvata in base al profilo. Se non è presente il profilo
allora viene mostrato il profilo di default.

### Invio email
Per inviare un'email è necessario eseguire il seguente comando:

```bash
./bin/go-mail-local [-profile=] sendmail
```
Dopo aver lanciato il comando verranno chiesti tutti i parametri necessari per inviare l'email.


