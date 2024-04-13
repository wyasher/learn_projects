package com.ruyuan.payment.server.response;

public class CommonResponse<T> {
    private boolean success = true;
    private String message;
    private  T content;


    public boolean isSuccess() {
        return success;
    }

    public void setSuccess(boolean success) {
        this.success = success;
    }

    public String getMessage() {
        return message;
    }

    public void setMessage(String message) {
        this.message = message;
    }

    public T getContent() {
        return content;
    }

    public void setContent(T content) {
        this.content = content;
    }
    @Override
    public String toString() {
        return "ResponseDto{" + "success=" + success +
                ", message='" + message + '\'' +
                ", content=" + content +
                '}';
    }

}
