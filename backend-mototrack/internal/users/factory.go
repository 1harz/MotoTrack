package users

import ("gitlab.com/prof-caio-de-melo/user-service/internal/utility"
		"encoding/json"
		"net/mail"
		"regexp"
		"errors")


type UserFactory struct {}
//type MotoFactory struct {}
func NewUserFactory() *UserFactory {
	return &UserFactory{}
}
func (f *UserFactory) CreateFromJSON(data []byte) (*User, error) {
	
	var user User
	
	if err := json.Unmarshal(data, &user); err != nil {
		return nil, err
	}
	
	return f.Validate(&user)
}

func (f *UserFactory) Validate(user *User) (*User, error) {
	
	/*if err := f.validateName(user); err != nil {
		return user, err
	}*/

	if err := f.validateEmail(user); err != nil {
		return user, err
	}

	if err := f.validatePassword(user); err != nil {
		return user, err
	}
	
	f.ensureID(user)
	f.ensureMetadata(user)
	for _, bike := range user.Bikes {
        if bike.MtID == "" {
            return nil, errors.New("each bike must have an ID")
        }
    }
	return user, nil
}


/* Validações para assegurar a criação de um usuário válido pela fábrica */

const (
	MIN_PASSWORD_SIZE = 8 // Tamanho mínimo de uma senha
	MAX_NUMBER_OF_METADATA_KEYS = 24 // Número máximo de metadados
	MAX_SIZE_OF_METADATA_VALUE = 512 // Tamanho máximo de um metadado, em bytes.
)

func (f *UserFactory) ensureID(user *User) {

	// 1. Verifica se o ID é uma string vazia.
	// Caso não seja, ALTERA o ID do usuário para um UUID.
	if user.ID == "" {
		user.ID = utility.GenerateStringID()
		return
	}

	// 2. Verifica se o ID representa um UUID.
	// Caso não seja, altera o ID do usuário para um UUID.
	if !utility.ValidateID(user.ID) {
		user.ID = utility.GenerateStringID()
		return
	}

	// 3. O id é válido, apenas finaliza sem alterações no usuário.
	return
} 

func (f *UserFactory) ensureMetadata(user *User) {
	
	// 1. Caso não tenha sido enviados metadados
	// Altere o usuário para ter metadados como um mapa vazio
	if user.Metadata == nil {
	 	user.Metadata = make(map[string]string)
	 	return
	}

	// 2. Caso os metadados tenham sido enviados,
	// assegure que os mesmos respeitem o tamanho máximo de chaves e valores.
	// Essa limitação imposta assegura que os metadados 
	// não sejam explorados por outros serviços para guardar informações gigantescas.
	// Os valores da limitação foram setados nas constantes: 
	// MAX_NUMBER_OF_METADATA_KEYS e MAX_SIZE_OF_METADATA_VALUE

	// 2.1 Caso o número de metadados enviados seja maior que o permitido,
	// os metadados são zerados, um mapa vazio.
	if len(user.Metadata) > MAX_NUMBER_OF_METADATA_KEYS {
		user.Metadata = make(map[string]string)
	 	return
	}

	// 2.2 Removendo metadados cujo os valores são maiores que MAX_SIZE_OF_METADATA_VALUE
	for key, value := range(user.Metadata) {
		if (len(value) > MAX_SIZE_OF_METADATA_VALUE) {
			delete(user.Metadata, key)
		}
	}

	return
}

/*func (f *UserFactory) validateName(user *User) error {
	// Em geral a validação de nomes não deve ser restrita.
	// A única coisa que irei fazer é assegurar que não é uma string vazia. 
	if user.Name == "" {
		return errors.New("invalid name: can not be empty")
	}

	return nil  
}*/

func (f *UserFactory) validateEmail(user *User) error {

	_, err := mail.ParseAddress(user.Email)
	return err
}
f
func (f *UserFactory) validatePassword(user *User) error {

	// Regra 1: A senha deve ter 8 caracteres, no minimo
	if len(user.Password) < MIN_PASSWORD_SIZE {
		return errors.New("invalid password: size is lower than " + string(MIN_PASSWORD_SIZE))
	}

	// Regra 2: A senha deve conter pelo menos um número
	if !regexp.MustCompile(`\d`).MatchString(user.Password) {
		return errors.New("invalid password: does not contain a number")
	}

	// Regra 3: A senha deve conter pelo menos uma letra maiuscula
	if !regexp.MustCompile(`.*[A-Z]`).MatchString(user.Password)  {
		return errors.New("invalid password: does not caontain a capital letter")
	}

	// Regra 4: A senha deve conter pelo menos um simbolo
	if !regexp.MustCompile(`[-+_!@#$%^&*.,?]`).MatchString(user.Password)  {
		return errors.New("invalid password: does not contain a valid simbol")
	}


	// Todas essas regras poderiam ser combinadas em um único regex.
	// deixei separado para facilitar o entendimento e clareza.
	return nil
}
