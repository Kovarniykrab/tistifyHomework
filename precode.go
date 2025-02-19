package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var cafeList = map[string][]string{
	"moscow": []string{"Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент"},
}

func mainHandle(w http.ResponseWriter, req *http.Request) {
	countStr := req.URL.Query().Get("count")
	if countStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("count missing"))
		return
	}

	count, err := strconv.Atoi(countStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("wrong count value"))
		return
	}

	city := req.URL.Query().Get("city")

	cafe, ok := cafeList[city]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("wrong city value"))
		return
	}

	if count > len(cafe) {
		count = len(cafe)
	}

	answer := strings.Join(cafe[:count], ",")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(answer))
}

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
