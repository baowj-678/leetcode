/**
 * @author: baowj
 * @time: 2022/8/21 17:18:04
 */
package main

func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	energyInc := 0
	experienceInc := 0
	for i := 0; i < len(energy); i++ {
		energyTmp, experienceTmp := energy[i], experience[i]
		if initialEnergy <= energyTmp {
			energyInc += energyTmp - initialEnergy + 1
			initialEnergy = energyTmp + 1
		}
		if initialExperience <= experienceTmp {
			experienceInc += experienceTmp - initialExperience + 1
			initialExperience = experienceTmp + 1
		}
		initialEnergy -= energyTmp
		initialExperience += experienceTmp
	}
	return energyInc + experienceInc
}
