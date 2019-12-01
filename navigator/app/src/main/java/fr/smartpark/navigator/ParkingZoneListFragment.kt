package fr.smartpark.navigator

import android.os.Bundle
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.fragment.app.Fragment
import androidx.fragment.app.viewModels
import androidx.lifecycle.observe
import fr.smartpark.navigator.adapters.ParkingZoneAdapter
import fr.smartpark.navigator.databinding.FragmentParkingZoneListBinding
import fr.smartpark.navigator.utilities.InjectorUtils
import fr.smartpark.navigator.viewmodels.ParkingZoneListViewModel

class ParkingZoneListFragment : Fragment() {
    private lateinit var binding: FragmentParkingZoneListBinding
    private val viewModel: ParkingZoneListViewModel by viewModels {
        InjectorUtils.provideParkingZoneListViewModelFactory(requireContext())
    }

    override fun onCreateView(
        inflater: LayoutInflater, container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        binding = FragmentParkingZoneListBinding.inflate(inflater, container, false)

        val adapter = ParkingZoneAdapter()
        binding.zoneList.adapter = adapter
        viewModel.zones.observe(viewLifecycleOwner) { zones ->
            adapter.submitList(zones)
        }
        return binding.root
    }
}
