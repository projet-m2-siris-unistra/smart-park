package fr.smartpark.navigator.utilities

import androidx.lifecycle.LiveData
import androidx.lifecycle.liveData
import androidx.lifecycle.map
import fr.smartpark.navigator.api.ApiResult
import kotlinx.coroutines.Dispatchers

fun <T, A> resultLiveData(databaseQuery: () -> LiveData<T>,
                          networkCall: suspend () -> ApiResult<A>,
                          saveCallResult: suspend (A) -> Unit): LiveData<ApiResult<T>> =
    liveData(Dispatchers.IO) {
        emit(ApiResult.loading())
        val source = databaseQuery.invoke().map { ApiResult.cached(it) }
        emitSource(source)

        val responseStatus = networkCall.invoke()
        if (responseStatus.status == ApiResult.Status.SUCCESS) {
            saveCallResult(responseStatus.data!!)
        } else if (responseStatus.status == ApiResult.Status.ERROR) {
            emit(ApiResult.error(responseStatus.message!!))
            emitSource(source)
        }
    }