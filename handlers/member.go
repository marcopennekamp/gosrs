package handlers
import (
	"github.com/marcopennekamp/gosrs/core"
	"net/http"
	"github.com/julienschmidt/httprouter"
	// "golang.org/x/crypto/bcrypt"
)

func ShowLogin(ctx *core.Context, w http.ResponseWriter, r *http.Request, _ httprouter.Params) (core.HttpStatus, error) {
	w.Write([]byte("GET login"))
	return http.StatusOK, nil
}

func Login(ctx *core.Context, w http.ResponseWriter, r *http.Request, _ httprouter.Params) (core.HttpStatus, error) {
	w.Write([]byte("POST login"))
	return http.StatusOK, nil
}


type InputFieldData struct {
	Value string
	Error string
}

func NewInputFieldData() InputFieldData {
	return InputFieldData{Value: "", Error: ""}
}

func (d *InputFieldData) HasError() bool {
	return d.Error != ""
}

type RegistrationFormData struct {
	Name 		InputFieldData
	Email		InputFieldData
	Password	InputFieldData
}

func NewRegistrationFormData() *RegistrationFormData {
	return &RegistrationFormData{Name: NewInputFieldData(), Email: NewInputFieldData(), Password: NewInputFieldData()}
}

func ShowRegistration(ctx *core.Context, w http.ResponseWriter, r *http.Request, _ httprouter.Params) (core.HttpStatus, error) {
	err := ctx.Templates.RenderTemplate("member/register", w, NewRegistrationFormData())
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func Register(ctx *core.Context, w http.ResponseWriter, r *http.Request, _ httprouter.Params) (core.HttpStatus, error) {
	name, email, password := r.PostFormValue("name"), r.PostFormValue("email"), r.PostFormValue("password")

	// Check whether the form is valid and display the form with errors if not.
	// TODO: Validate the form with regular expressions.
	valid := name != "" && email != "" && password != ""
	if !valid {
		fd := NewRegistrationFormData()
		fd.Name.Value = name
		fd.Email.Value = email
		fd.Password.Value = password
		if name == "" {
			fd.Name.Value = "" // Remove the value if it's invalid.
			fd.Name.Error = "The username must not be empty."
		}
		if email == "" {
			fd.Email.Value = ""
			fd.Email.Error = "The email must not be empty."
		}
		if password == "" {
			fd.Password.Value = ""
			fd.Password.Error = "The password must not be empty."
		}

		err := ctx.Templates.RenderTemplate("member/register", w, fd)
		if err != nil {
			return http.StatusInternalServerError, err
		}

		return http.StatusOK, nil
	}

	// Check if the username or email are already taken.
	rows, err := ctx.Db.NamedQuery(
		`SELECT COUNT(*) FROM member WHERE name = :name OR email = :email`,
		map[string]interface{} {
			"name": name,
			"email": email,
		},
	)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	for rows.Next() {
		cols, err := rows.SliceScan()
		if err != nil {
			return http.StatusInternalServerError, err
		}
		for _, col := range cols {
			println(col)
		}
	}

	w.Write([]byte("POST register"))
	return http.StatusOK, nil
}
