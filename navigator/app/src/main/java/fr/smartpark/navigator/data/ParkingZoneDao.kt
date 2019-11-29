package fr.smartpark.navigator.data

import androidx.lifecycle.LiveData
import androidx.room.Dao
import androidx.room.Insert
import androidx.room.Query

@Dao
interface ParkingZoneDao {
    @Query("SELECT * FROM zones ORDER BY id")
    fun getZones(): LiveData<List<ParkingZone>>

    @Insert
    suspend fun insertAll(zones: List<ParkingZone>)
}