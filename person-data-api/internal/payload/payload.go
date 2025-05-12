package payload

// Структура для данных о человеке(от 2 символов, поля обязательные, только буквы)
type PersonPayload struct {
	Name       string `json:"name" validate:"required,min=2,alpha"`
	Surname    string `json:"surname" validate:"required,min=2,alpha"`
	Patronymic string `json:"patronymic,omitempty" validate:"omitempty,min=2,alpha"`
}
