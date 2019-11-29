package fr.smartpark.navigator.workers

import android.content.Context
import android.util.Log
import androidx.work.CoroutineWorker
import androidx.work.WorkerParameters
import com.google.gson.Gson
import com.google.gson.reflect.TypeToken
import com.google.gson.stream.JsonReader
import fr.smartpark.navigator.data.AppDatabase
import fr.smartpark.navigator.data.ParkingZone
import fr.smartpark.navigator.utilities.PARKING_DATA_FILE
import kotlinx.coroutines.coroutineScope

class SeedDatabaseWorker(
    context: Context,
    workerParams: WorkerParameters
) : CoroutineWorker(context, workerParams) {
    override suspend fun doWork(): Result = coroutineScope {
        try {
            applicationContext.assets.open(PARKING_DATA_FILE).use { inputStream ->
                JsonReader(inputStream.reader()).use { jsonReader ->
                    val zoneType = object : TypeToken<List<ParkingZone>>() {}.type
                    val zoneList: List<ParkingZone> = Gson().fromJson(jsonReader, zoneType)

                    val database = AppDatabase(applicationContext)
                    database.zonesDao().insertAll(zoneList)

                    Result.success()
                }
            }
        } catch (ex: Exception) {
            Log.e(TAG, "Error seeding database", ex)
            Result.failure()
        }
    }

    companion object {
        private val TAG = SeedDatabaseWorker::class.java.simpleName
    }
}
