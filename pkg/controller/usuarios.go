package controller


//Usuarios esto es una estructura
type Usuarios struct {
    IDUsuario int64  `json:"id_usuario"`
    Nombre    string `json:"nombre"`
    Usuario   string `json:"usuario"`
    Password  string `json:"password"`
}

//LoginUsuario estructura para validar el login
type LoginUsuario struct {
    Usuario  string `json:"usuario"`
    Password string `json:"password"`
}

func ParseLoginUsuario(data []byte) (*usuarios.LoginUsuario, error) {
    var usuario usuarios.LoginUsuario
    if err := json.Unmarshal(data, &usuario); err != nil {
        return nil, err
    }

    return &usuario, nil
}

func ParseUsuario(data []byte) (*usuarios.Usuarios, error) {
    var usuario usuarios.Usuarios
    if err := json.Unmarshal(data, &usuario); err != nil {
        return nil, err
    }

    return &usuario, nil
}

//CreateUsuario crea un usuario
func CreateUsuario(usuario *usuarios.Usuarios) *apierror.ApiError {
	stmt, err := db.Init().Prepare("insert into usuarios (nombre, usuario, password) values(?,?,?);")

    if err != nil {
		return &apierror.ApiError {
			Status: http.StatusInternalServerError,
			Message: "Data base error",
		}
    }

	_, err = stmt.Exec(usuario.Nombre, usuario.Usuario, usuario.Password)

	if err != nil {
		return &apierror.ApiError {
			Status: http.StatusInternalServerError,
			Message: "Error while saving the username data",
		}
	}

    defer stmt.Close()
    return nil
}

//UpdateUsuario modifica el nombre y la password del usuario
func UpdateUsuario(user int64, usuario *usuarios.Usuarios) *apierror.ApiError {
	stmt, err := db.Init().Prepare("update usuarios set nombre=?, password=? where id_usuario=?;")

    if err != nil {
        return &apierror.ApiError {
			Status: http.StatusInternalServerError,
			Message: "Data base error",
		}
    }

	_, err = stmt.Exec(usuario.Nombre, usuario.Password, user)

	if err != nil {
		return &apierror.ApiError {
			Status: http.StatusInternalServerError,
			Message: "Error while updating the username data",
		}
	}

    defer stmt.Close()
    return nil
}

//Login comprueba que el usuario esta registrado en la pagina
func Login(usuario *usuarios.LoginUsuario) (*usuarios.Usuarios, *apierror.ApiError) {
    var user usuarios.Usuarios
    stmt, err := db.Init().Prepare("select * from usuarios where usuario = ? and password = ?;")

    if err != nil {
		return nil, &apierror.ApiError {
			Status: http.StatusInternalServerError,
			Message: "Data base error",
		}
    }

	result := stmt.QueryRow(usuario.Usuario, usuario.Password)
	err = result.Scan(
		&user.IDUsuario,
        &user.Nombre,
        &user.Usuario,
		&user.Password)

	if err != nil {
		return nil, &apierror.ApiError {
			Status: http.StatusUnauthorized,
			Message: "Wrong username or password",
		}
	}

	defer stmt.Close()
    return &user, nil
}

//Logout elimina la sesion del usuario y lo desloguea
func Logout(user int64) *apierror.ApiError {
    stmt, err := db.Init().Prepare("delete from sessions where user = ?;")

    if err != nil {
        return &apierror.ApiError {
			Status: http.StatusInternalServerError,
			Message: "Data base error",
		}
	}

	_, err = stmt.Exec(user)

	if err != nil {
		return &apierror.ApiError {
			Status: http.StatusInternalServerError,
			Message: "Data base error",
		}
	}

    defer stmt.Close()
    return nil
}
