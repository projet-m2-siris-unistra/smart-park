package fr.smartpark.navigator

import android.os.Bundle
import androidx.appcompat.app.AppCompatActivity
import androidx.databinding.DataBindingUtil
import fr.smartpark.navigator.databinding.ActivityZonesBinding

class ZonesActivity : AppCompatActivity() {
    private lateinit var binding: ActivityZonesBinding

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        binding = DataBindingUtil.setContentView(this, R.layout.activity_zones)
    }
}
