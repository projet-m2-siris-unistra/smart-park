package fr.smartpark.navigator.data

import androidx.room.ColumnInfo
import androidx.room.Entity
import androidx.room.PrimaryKey

@Entity(tableName = "zones")
data class ParkingZone(
    @PrimaryKey @ColumnInfo(name = "id") val zoneId: String,
    val name: String
)