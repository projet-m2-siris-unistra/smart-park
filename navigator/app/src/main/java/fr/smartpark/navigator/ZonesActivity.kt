package fr.smartpark.navigator

import android.os.Bundle
import android.util.Log
import androidx.appcompat.app.AppCompatActivity
import androidx.databinding.DataBindingUtil
import androidx.navigation.navArgs
import fr.smartpark.navigator.databinding.ActivityZonesBinding

class ZonesActivity : AppCompatActivity() {
    private lateinit var binding: ActivityZonesBinding
    private val args: ZonesActivityArgs by navArgs()

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        binding = DataBindingUtil.setContentView(this, R.layout.activity_zones)
        Log.d("ZoneActivity", args.tenantId.toString())
    }
}
