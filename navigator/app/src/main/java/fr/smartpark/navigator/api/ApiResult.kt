package fr.smartpark.navigator.api


data class ApiResult<out T>(val status: Status, val data: T?, val message: String?) {

    enum class Status {
        SUCCESS,
        CACHED,
        ERROR,
        LOADING
    }

    fun <R> map(mapper: (it: T) -> R): ApiResult<R> =
        when(status) {
            Status.SUCCESS -> ApiResult(Status.SUCCESS, mapper(data!!), message)
            Status.CACHED -> ApiResult(Status.SUCCESS, mapper(data!!), message)
            else -> ApiResult(status, null, message)
        }

    companion object {
        fun <T> cached(data: T): ApiResult<T> {
            return ApiResult(
                Status.CACHED,
                data,
                null
            )
        }

        fun <T> success(data: T): ApiResult<T> {
            return ApiResult(
                Status.SUCCESS,
                data,
                null
            )
        }

        fun <T> error(message: String, data: T? = null): ApiResult<T> {
            return ApiResult(
                Status.ERROR,
                data,
                message
            )
        }

        fun <T> loading(data: T? = null): ApiResult<T> {
            return ApiResult(
                Status.LOADING,
                data,
                null
            )
        }
    }
}