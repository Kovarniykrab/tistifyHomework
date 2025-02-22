package tistifyhomework

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	// здесь нужно добавить необходимые проверки

	status := responseRecorder.Code
	assert.Equal(t, status, http.StatusOK)

	body := responseRecorder.Body.String()
	assert.NotEmpty(t, body, "тело запроса пустое")

	list := strings.Split(body, ",")

	city := req.URL.Query().Get("city")
	expectedLenght := len(cafeList[city])
	assert.Len(t, list, expectedLenght)
	//я знаю, что я убрал  totalCount :=4 и было бы проще написать код такого рода:
	// assert.Len(t, list, totalCount), но мне было очень интересно попробовать решить задачу таким образом
	//и я считаю, что так будет правильнее. Ведь пользователь может ввести любой запрос
	//условно я представил, что в мапе лежит больше одного города
	//надеюсь Вы простите меня за небольшую шалось)
	//еще у меня почему-то тесты не запускаются, поэтому надеюсь, что все должно быть правильно
}

func TestMainHandlerNice(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	// здесь нужно добавить необходимые проверки

	status := responseRecorder.Code
	assert.Equal(t, status, http.StatusOK)

	body := responseRecorder.Body.String()
	assert.NotEmpty(t, body, "тело запроса пустое")

	assert.Equal(t, body, "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент")
}

func TestMainHandlerBadRequest(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=1&city=Tomsk", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	// здесь нужно добавить необходимые проверки

	status := responseRecorder.Code
	assert.Equal(t, status, http.StatusBadRequest)

	body := responseRecorder.Body.String()
	assert.Equal(t, "wrong city value", body)

}
