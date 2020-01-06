package fr.smartpark.navigator.data

import androidx.lifecycle.LiveData
import androidx.room.Dao
import androidx.room.Insert
import androidx.room.Query
import fr.smartpark.navigator.data.models.Zone

@Dao
interface ZoneDao {
    @Query("SELECT * FROM zones ORDER BY id")
    fun getZones(): LiveData<List<Zone>>

    @Query("SELECT * from zones WHERE id = :id")
    fun getZone(id: Long): LiveData<Zone>

    @Insert
    suspend fun insertAll(zones: List<Zone>)
}