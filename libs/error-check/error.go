package errorcheck

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

// CustomError ...
type CustomError struct {
	ErrorCode    int64  `json:"errorCode"`    // 에러 코드
	ErrorMessage string `json:"errorMessage"` // 에러 메세지
	Message      string `json:"message"`      // 클라이언트에서 확인하는 에러 메세지
	StatusCode   int32  `json:"statusCode"`   // http 상태코드
	CreatedAt    int64  `json:"createdAt"`    // 생성날짜
}

// ERROR_CODE_MAP 에러코드 맵
var ERROR_CODE_MAP = map[int64]string{
	// Undefined Error
	00000: "서버 오류",
	01000: "DB 문법에러",
	01001: "DB 에러",
	01002: "입력값이 잘못되었습니다.",
	// Account service error
	10001: "이미 존재하는 아이디입니다.",
	10002: "이미 존재하는 이메일입니다.",
	10003: "패스워드가 일치하지 않습니다.",
	10004: "존재하지 않는 아이디입니다.",
	10005: "존재하지 않는 이메일입니다.",
	10006: "코드값이 정확하지 않습니다.",
	10007: "이메일이 형식이 유효하지 않습니다.",
	10008: "인증되지 않은 이메일입니다.",
	10009: "회원 가입에 실패했습니다.",
	10010: "인증이 잘못되었습니다.",
	10011: "다시 시도해 주세요.",
	10012: "비밀번호가 취약합니다.",
	10013: "5분 후에 다시 시도해 주세요.",
	10014: "관리자 승인 후 이용해 주세요.",
	10015: "학교 정보를 입력해 주세요.",
	10016: "인증 코드 오류입니다.",
	20001: "인증이 잘못되었습니다.",
	20002: "잘못된 접근입니다.",
	20003: "코드값이 정확하지 않습니다.",
	20004: "권한이 없습니다.",
}

// Error 에러리턴
func Error(errorCode int, err error) error {
	_, exists := ERROR_CODE_MAP[int64(errorCode)]
	if !exists {
		errorCode = 00000
	}

	customError := CustomError{
		ErrorCode:    int64(errorCode),
		ErrorMessage: err.Error(),
		Message:      ERROR_CODE_MAP[int64(errorCode)],
		StatusCode:   http.StatusServiceUnavailable,
		CreatedAt:    time.Now().Local().Unix(),
	}
	jsonBytes, _ := json.Marshal(customError)
	return errors.New(string(jsonBytes))
}
