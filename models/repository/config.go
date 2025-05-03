package repository

type Config struct {
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	PostgresSslMode  string
	JwtSecret        string
	JwtExpiration    string
	AiService        string
	FrontendHost     string
	Port             string
	BrevoKey         string
	AppEnv           string
	MailSender       string
	MailSmtpHost     string
	MailSmtpUsername string
	MailSmtpPassword string
}
