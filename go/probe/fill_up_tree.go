package probe

import (
	"fmt"
	"sort"
	"strings"

	gongtree_buttons "github.com/fullstack-lang/gongtree/go/buttons"
	tree "github.com/fullstack-lang/gongtree/go/models"

	gong_models "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gongmusicxml/go/models"
)

func fillUpTree(
	probe *Probe,
) {
	// keep in memory which nodes have been unfolded / folded
	expandedNodesSet := make(map[string]any, 0)
	var _sidebar *tree.Tree
	for __sidebar := range probe.treeStage.Trees {
		_sidebar = __sidebar
	}
	if _sidebar != nil {
		for _, node := range _sidebar.RootNodes {
			if node.IsExpanded {
				expandedNodesSet[strings.Fields(node.Name)[0]] = true
			}
		}
	}

	probe.treeStage.Reset()

	// create tree
	sidebar := (&tree.Tree{Name: "gong"}).Stage(probe.treeStage)

	// collect all gong struct to construe the true
	setOfGongStructs := *gong_models.GetGongstructInstancesSet[gong_models.GongStruct](probe.gongStage)

	sliceOfGongStructsSorted := make([]*gong_models.GongStruct, len(setOfGongStructs))
	i := 0
	for k := range setOfGongStructs {
		sliceOfGongStructsSorted[i] = k
		i++
	}
	sort.Slice(sliceOfGongStructsSorted, func(i, j int) bool {
		return sliceOfGongStructsSorted[i].Name < sliceOfGongStructsSorted[j].Name
	})

	for _, gongStruct := range sliceOfGongStructsSorted {

		name := gongStruct.Name + " (" +
			fmt.Sprintf("%d", probe.stageOfInterest.Map_GongStructName_InstancesNb[gongStruct.Name]) + ")"

		nodeGongstruct := (&tree.Node{Name: name}).Stage(probe.treeStage)

		nodeGongstruct.IsExpanded = false
		if _, ok := expandedNodesSet[strings.Fields(name)[0]]; ok {
			nodeGongstruct.IsExpanded = true
		}

		switch gongStruct.Name {
		// insertion point
		case "Accidental":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Accidental](probe.stageOfInterest)
			for _accidental := range set {
				nodeInstance := (&tree.Node{Name: _accidental.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_accidental, "Accidental", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Accidental_mark":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Accidental_mark](probe.stageOfInterest)
			for _accidental_mark := range set {
				nodeInstance := (&tree.Node{Name: _accidental_mark.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_accidental_mark, "Accidental_mark", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Accidental_text":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Accidental_text](probe.stageOfInterest)
			for _accidental_text := range set {
				nodeInstance := (&tree.Node{Name: _accidental_text.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_accidental_text, "Accidental_text", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Accord":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Accord](probe.stageOfInterest)
			for _accord := range set {
				nodeInstance := (&tree.Node{Name: _accord.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_accord, "Accord", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Accordion_registration":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Accordion_registration](probe.stageOfInterest)
			for _accordion_registration := range set {
				nodeInstance := (&tree.Node{Name: _accordion_registration.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_accordion_registration, "Accordion_registration", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "AnyType":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.AnyType](probe.stageOfInterest)
			for _anytype := range set {
				nodeInstance := (&tree.Node{Name: _anytype.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_anytype, "AnyType", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Appearance":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Appearance](probe.stageOfInterest)
			for _appearance := range set {
				nodeInstance := (&tree.Node{Name: _appearance.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_appearance, "Appearance", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Arpeggiate":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Arpeggiate](probe.stageOfInterest)
			for _arpeggiate := range set {
				nodeInstance := (&tree.Node{Name: _arpeggiate.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_arpeggiate, "Arpeggiate", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Arrow":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Arrow](probe.stageOfInterest)
			for _arrow := range set {
				nodeInstance := (&tree.Node{Name: _arrow.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_arrow, "Arrow", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Articulations":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Articulations](probe.stageOfInterest)
			for _articulations := range set {
				nodeInstance := (&tree.Node{Name: _articulations.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_articulations, "Articulations", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Assess":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Assess](probe.stageOfInterest)
			for _assess := range set {
				nodeInstance := (&tree.Node{Name: _assess.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_assess, "Assess", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Attributes":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Attributes](probe.stageOfInterest)
			for _attributes := range set {
				nodeInstance := (&tree.Node{Name: _attributes.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attributes, "Attributes", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Backup":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Backup](probe.stageOfInterest)
			for _backup := range set {
				nodeInstance := (&tree.Node{Name: _backup.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_backup, "Backup", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Bar_style_color":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Bar_style_color](probe.stageOfInterest)
			for _bar_style_color := range set {
				nodeInstance := (&tree.Node{Name: _bar_style_color.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_bar_style_color, "Bar_style_color", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Barline":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Barline](probe.stageOfInterest)
			for _barline := range set {
				nodeInstance := (&tree.Node{Name: _barline.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_barline, "Barline", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Barre":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Barre](probe.stageOfInterest)
			for _barre := range set {
				nodeInstance := (&tree.Node{Name: _barre.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_barre, "Barre", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Bass":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Bass](probe.stageOfInterest)
			for _bass := range set {
				nodeInstance := (&tree.Node{Name: _bass.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_bass, "Bass", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Bass_step":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Bass_step](probe.stageOfInterest)
			for _bass_step := range set {
				nodeInstance := (&tree.Node{Name: _bass_step.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_bass_step, "Bass_step", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Beam":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Beam](probe.stageOfInterest)
			for _beam := range set {
				nodeInstance := (&tree.Node{Name: _beam.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_beam, "Beam", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Beat_repeat":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Beat_repeat](probe.stageOfInterest)
			for _beat_repeat := range set {
				nodeInstance := (&tree.Node{Name: _beat_repeat.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_beat_repeat, "Beat_repeat", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Beat_unit_tied":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Beat_unit_tied](probe.stageOfInterest)
			for _beat_unit_tied := range set {
				nodeInstance := (&tree.Node{Name: _beat_unit_tied.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_beat_unit_tied, "Beat_unit_tied", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Beater":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Beater](probe.stageOfInterest)
			for _beater := range set {
				nodeInstance := (&tree.Node{Name: _beater.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_beater, "Beater", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Bend":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Bend](probe.stageOfInterest)
			for _bend := range set {
				nodeInstance := (&tree.Node{Name: _bend.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_bend, "Bend", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Bookmark":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Bookmark](probe.stageOfInterest)
			for _bookmark := range set {
				nodeInstance := (&tree.Node{Name: _bookmark.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_bookmark, "Bookmark", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Bracket":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Bracket](probe.stageOfInterest)
			for _bracket := range set {
				nodeInstance := (&tree.Node{Name: _bracket.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_bracket, "Bracket", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Breath_mark":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Breath_mark](probe.stageOfInterest)
			for _breath_mark := range set {
				nodeInstance := (&tree.Node{Name: _breath_mark.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_breath_mark, "Breath_mark", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Caesura":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Caesura](probe.stageOfInterest)
			for _caesura := range set {
				nodeInstance := (&tree.Node{Name: _caesura.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_caesura, "Caesura", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Cancel":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Cancel](probe.stageOfInterest)
			for _cancel := range set {
				nodeInstance := (&tree.Node{Name: _cancel.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_cancel, "Cancel", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Clef":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Clef](probe.stageOfInterest)
			for _clef := range set {
				nodeInstance := (&tree.Node{Name: _clef.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_clef, "Clef", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Coda":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Coda](probe.stageOfInterest)
			for _coda := range set {
				nodeInstance := (&tree.Node{Name: _coda.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_coda, "Coda", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Credit":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Credit](probe.stageOfInterest)
			for _credit := range set {
				nodeInstance := (&tree.Node{Name: _credit.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_credit, "Credit", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Dashes":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Dashes](probe.stageOfInterest)
			for _dashes := range set {
				nodeInstance := (&tree.Node{Name: _dashes.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_dashes, "Dashes", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Defaults":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Defaults](probe.stageOfInterest)
			for _defaults := range set {
				nodeInstance := (&tree.Node{Name: _defaults.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_defaults, "Defaults", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Degree":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Degree](probe.stageOfInterest)
			for _degree := range set {
				nodeInstance := (&tree.Node{Name: _degree.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_degree, "Degree", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Degree_alter":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Degree_alter](probe.stageOfInterest)
			for _degree_alter := range set {
				nodeInstance := (&tree.Node{Name: _degree_alter.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_degree_alter, "Degree_alter", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Degree_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Degree_type](probe.stageOfInterest)
			for _degree_type := range set {
				nodeInstance := (&tree.Node{Name: _degree_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_degree_type, "Degree_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Degree_value":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Degree_value](probe.stageOfInterest)
			for _degree_value := range set {
				nodeInstance := (&tree.Node{Name: _degree_value.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_degree_value, "Degree_value", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Direction":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Direction](probe.stageOfInterest)
			for _direction := range set {
				nodeInstance := (&tree.Node{Name: _direction.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_direction, "Direction", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Direction_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Direction_type](probe.stageOfInterest)
			for _direction_type := range set {
				nodeInstance := (&tree.Node{Name: _direction_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_direction_type, "Direction_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Distance":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Distance](probe.stageOfInterest)
			for _distance := range set {
				nodeInstance := (&tree.Node{Name: _distance.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_distance, "Distance", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Double":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Double](probe.stageOfInterest)
			for _double := range set {
				nodeInstance := (&tree.Node{Name: _double.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_double, "Double", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Dynamics":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Dynamics](probe.stageOfInterest)
			for _dynamics := range set {
				nodeInstance := (&tree.Node{Name: _dynamics.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_dynamics, "Dynamics", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Effect":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Effect](probe.stageOfInterest)
			for _effect := range set {
				nodeInstance := (&tree.Node{Name: _effect.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_effect, "Effect", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Elision":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Elision](probe.stageOfInterest)
			for _elision := range set {
				nodeInstance := (&tree.Node{Name: _elision.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_elision, "Elision", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Empty":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Empty](probe.stageOfInterest)
			for _empty := range set {
				nodeInstance := (&tree.Node{Name: _empty.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_empty, "Empty", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Empty_font":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Empty_font](probe.stageOfInterest)
			for _empty_font := range set {
				nodeInstance := (&tree.Node{Name: _empty_font.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_empty_font, "Empty_font", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Empty_line":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Empty_line](probe.stageOfInterest)
			for _empty_line := range set {
				nodeInstance := (&tree.Node{Name: _empty_line.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_empty_line, "Empty_line", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Empty_placement":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Empty_placement](probe.stageOfInterest)
			for _empty_placement := range set {
				nodeInstance := (&tree.Node{Name: _empty_placement.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_empty_placement, "Empty_placement", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Empty_placement_smufl":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Empty_placement_smufl](probe.stageOfInterest)
			for _empty_placement_smufl := range set {
				nodeInstance := (&tree.Node{Name: _empty_placement_smufl.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_empty_placement_smufl, "Empty_placement_smufl", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Empty_print_object_style_align":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Empty_print_object_style_align](probe.stageOfInterest)
			for _empty_print_object_style_align := range set {
				nodeInstance := (&tree.Node{Name: _empty_print_object_style_align.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_empty_print_object_style_align, "Empty_print_object_style_align", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Empty_print_style":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Empty_print_style](probe.stageOfInterest)
			for _empty_print_style := range set {
				nodeInstance := (&tree.Node{Name: _empty_print_style.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_empty_print_style, "Empty_print_style", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Empty_print_style_align":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Empty_print_style_align](probe.stageOfInterest)
			for _empty_print_style_align := range set {
				nodeInstance := (&tree.Node{Name: _empty_print_style_align.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_empty_print_style_align, "Empty_print_style_align", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Empty_print_style_align_id":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Empty_print_style_align_id](probe.stageOfInterest)
			for _empty_print_style_align_id := range set {
				nodeInstance := (&tree.Node{Name: _empty_print_style_align_id.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_empty_print_style_align_id, "Empty_print_style_align_id", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Empty_trill_sound":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Empty_trill_sound](probe.stageOfInterest)
			for _empty_trill_sound := range set {
				nodeInstance := (&tree.Node{Name: _empty_trill_sound.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_empty_trill_sound, "Empty_trill_sound", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Encoding":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Encoding](probe.stageOfInterest)
			for _encoding := range set {
				nodeInstance := (&tree.Node{Name: _encoding.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_encoding, "Encoding", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Ending":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Ending](probe.stageOfInterest)
			for _ending := range set {
				nodeInstance := (&tree.Node{Name: _ending.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_ending, "Ending", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Extend":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Extend](probe.stageOfInterest)
			for _extend := range set {
				nodeInstance := (&tree.Node{Name: _extend.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_extend, "Extend", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Feature":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Feature](probe.stageOfInterest)
			for _feature := range set {
				nodeInstance := (&tree.Node{Name: _feature.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_feature, "Feature", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Fermata":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Fermata](probe.stageOfInterest)
			for _fermata := range set {
				nodeInstance := (&tree.Node{Name: _fermata.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_fermata, "Fermata", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Figure":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Figure](probe.stageOfInterest)
			for _figure := range set {
				nodeInstance := (&tree.Node{Name: _figure.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_figure, "Figure", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Figured_bass":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Figured_bass](probe.stageOfInterest)
			for _figured_bass := range set {
				nodeInstance := (&tree.Node{Name: _figured_bass.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_figured_bass, "Figured_bass", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Fingering":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Fingering](probe.stageOfInterest)
			for _fingering := range set {
				nodeInstance := (&tree.Node{Name: _fingering.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_fingering, "Fingering", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "First_fret":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.First_fret](probe.stageOfInterest)
			for _first_fret := range set {
				nodeInstance := (&tree.Node{Name: _first_fret.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_first_fret, "First_fret", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Foo":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Foo](probe.stageOfInterest)
			for _foo := range set {
				nodeInstance := (&tree.Node{Name: _foo.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_foo, "Foo", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "For_part":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.For_part](probe.stageOfInterest)
			for _for_part := range set {
				nodeInstance := (&tree.Node{Name: _for_part.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_for_part, "For_part", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Formatted_symbol":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Formatted_symbol](probe.stageOfInterest)
			for _formatted_symbol := range set {
				nodeInstance := (&tree.Node{Name: _formatted_symbol.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_formatted_symbol, "Formatted_symbol", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Formatted_symbol_id":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Formatted_symbol_id](probe.stageOfInterest)
			for _formatted_symbol_id := range set {
				nodeInstance := (&tree.Node{Name: _formatted_symbol_id.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_formatted_symbol_id, "Formatted_symbol_id", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Forward":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Forward](probe.stageOfInterest)
			for _forward := range set {
				nodeInstance := (&tree.Node{Name: _forward.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_forward, "Forward", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Frame":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Frame](probe.stageOfInterest)
			for _frame := range set {
				nodeInstance := (&tree.Node{Name: _frame.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_frame, "Frame", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Frame_note":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Frame_note](probe.stageOfInterest)
			for _frame_note := range set {
				nodeInstance := (&tree.Node{Name: _frame_note.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_frame_note, "Frame_note", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Fret":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Fret](probe.stageOfInterest)
			for _fret := range set {
				nodeInstance := (&tree.Node{Name: _fret.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_fret, "Fret", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Glass":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Glass](probe.stageOfInterest)
			for _glass := range set {
				nodeInstance := (&tree.Node{Name: _glass.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_glass, "Glass", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Glissando":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Glissando](probe.stageOfInterest)
			for _glissando := range set {
				nodeInstance := (&tree.Node{Name: _glissando.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_glissando, "Glissando", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Glyph":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Glyph](probe.stageOfInterest)
			for _glyph := range set {
				nodeInstance := (&tree.Node{Name: _glyph.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_glyph, "Glyph", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Grace":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Grace](probe.stageOfInterest)
			for _grace := range set {
				nodeInstance := (&tree.Node{Name: _grace.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_grace, "Grace", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Group_barline":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Group_barline](probe.stageOfInterest)
			for _group_barline := range set {
				nodeInstance := (&tree.Node{Name: _group_barline.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_group_barline, "Group_barline", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Group_symbol":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Group_symbol](probe.stageOfInterest)
			for _group_symbol := range set {
				nodeInstance := (&tree.Node{Name: _group_symbol.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_group_symbol, "Group_symbol", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Grouping":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Grouping](probe.stageOfInterest)
			for _grouping := range set {
				nodeInstance := (&tree.Node{Name: _grouping.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_grouping, "Grouping", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Hammer_on_pull_off":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Hammer_on_pull_off](probe.stageOfInterest)
			for _hammer_on_pull_off := range set {
				nodeInstance := (&tree.Node{Name: _hammer_on_pull_off.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_hammer_on_pull_off, "Hammer_on_pull_off", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Handbell":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Handbell](probe.stageOfInterest)
			for _handbell := range set {
				nodeInstance := (&tree.Node{Name: _handbell.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_handbell, "Handbell", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Harmon_closed":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Harmon_closed](probe.stageOfInterest)
			for _harmon_closed := range set {
				nodeInstance := (&tree.Node{Name: _harmon_closed.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_harmon_closed, "Harmon_closed", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Harmon_mute":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Harmon_mute](probe.stageOfInterest)
			for _harmon_mute := range set {
				nodeInstance := (&tree.Node{Name: _harmon_mute.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_harmon_mute, "Harmon_mute", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Harmonic":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Harmonic](probe.stageOfInterest)
			for _harmonic := range set {
				nodeInstance := (&tree.Node{Name: _harmonic.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_harmonic, "Harmonic", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Harmony":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Harmony](probe.stageOfInterest)
			for _harmony := range set {
				nodeInstance := (&tree.Node{Name: _harmony.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_harmony, "Harmony", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Harmony_alter":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Harmony_alter](probe.stageOfInterest)
			for _harmony_alter := range set {
				nodeInstance := (&tree.Node{Name: _harmony_alter.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_harmony_alter, "Harmony_alter", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Harp_pedals":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Harp_pedals](probe.stageOfInterest)
			for _harp_pedals := range set {
				nodeInstance := (&tree.Node{Name: _harp_pedals.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_harp_pedals, "Harp_pedals", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Heel_toe":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Heel_toe](probe.stageOfInterest)
			for _heel_toe := range set {
				nodeInstance := (&tree.Node{Name: _heel_toe.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_heel_toe, "Heel_toe", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Hole":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Hole](probe.stageOfInterest)
			for _hole := range set {
				nodeInstance := (&tree.Node{Name: _hole.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_hole, "Hole", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Hole_closed":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Hole_closed](probe.stageOfInterest)
			for _hole_closed := range set {
				nodeInstance := (&tree.Node{Name: _hole_closed.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_hole_closed, "Hole_closed", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Horizontal_turn":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Horizontal_turn](probe.stageOfInterest)
			for _horizontal_turn := range set {
				nodeInstance := (&tree.Node{Name: _horizontal_turn.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_horizontal_turn, "Horizontal_turn", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Identification":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Identification](probe.stageOfInterest)
			for _identification := range set {
				nodeInstance := (&tree.Node{Name: _identification.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_identification, "Identification", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Image":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Image](probe.stageOfInterest)
			for _image := range set {
				nodeInstance := (&tree.Node{Name: _image.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_image, "Image", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Instrument":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Instrument](probe.stageOfInterest)
			for _instrument := range set {
				nodeInstance := (&tree.Node{Name: _instrument.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_instrument, "Instrument", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Instrument_change":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Instrument_change](probe.stageOfInterest)
			for _instrument_change := range set {
				nodeInstance := (&tree.Node{Name: _instrument_change.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_instrument_change, "Instrument_change", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Instrument_link":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Instrument_link](probe.stageOfInterest)
			for _instrument_link := range set {
				nodeInstance := (&tree.Node{Name: _instrument_link.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_instrument_link, "Instrument_link", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Interchangeable":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Interchangeable](probe.stageOfInterest)
			for _interchangeable := range set {
				nodeInstance := (&tree.Node{Name: _interchangeable.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_interchangeable, "Interchangeable", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Inversion":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Inversion](probe.stageOfInterest)
			for _inversion := range set {
				nodeInstance := (&tree.Node{Name: _inversion.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_inversion, "Inversion", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Key":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Key](probe.stageOfInterest)
			for _key := range set {
				nodeInstance := (&tree.Node{Name: _key.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_key, "Key", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Key_accidental":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Key_accidental](probe.stageOfInterest)
			for _key_accidental := range set {
				nodeInstance := (&tree.Node{Name: _key_accidental.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_key_accidental, "Key_accidental", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Key_octave":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Key_octave](probe.stageOfInterest)
			for _key_octave := range set {
				nodeInstance := (&tree.Node{Name: _key_octave.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_key_octave, "Key_octave", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Kind":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Kind](probe.stageOfInterest)
			for _kind := range set {
				nodeInstance := (&tree.Node{Name: _kind.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_kind, "Kind", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Level":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Level](probe.stageOfInterest)
			for _level := range set {
				nodeInstance := (&tree.Node{Name: _level.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_level, "Level", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Line_detail":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Line_detail](probe.stageOfInterest)
			for _line_detail := range set {
				nodeInstance := (&tree.Node{Name: _line_detail.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_line_detail, "Line_detail", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Line_width":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Line_width](probe.stageOfInterest)
			for _line_width := range set {
				nodeInstance := (&tree.Node{Name: _line_width.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_line_width, "Line_width", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Link":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Link](probe.stageOfInterest)
			for _link := range set {
				nodeInstance := (&tree.Node{Name: _link.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_link, "Link", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Listen":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Listen](probe.stageOfInterest)
			for _listen := range set {
				nodeInstance := (&tree.Node{Name: _listen.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_listen, "Listen", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Listening":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Listening](probe.stageOfInterest)
			for _listening := range set {
				nodeInstance := (&tree.Node{Name: _listening.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_listening, "Listening", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Lyric":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Lyric](probe.stageOfInterest)
			for _lyric := range set {
				nodeInstance := (&tree.Node{Name: _lyric.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_lyric, "Lyric", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Lyric_font":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Lyric_font](probe.stageOfInterest)
			for _lyric_font := range set {
				nodeInstance := (&tree.Node{Name: _lyric_font.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_lyric_font, "Lyric_font", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Lyric_language":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Lyric_language](probe.stageOfInterest)
			for _lyric_language := range set {
				nodeInstance := (&tree.Node{Name: _lyric_language.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_lyric_language, "Lyric_language", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Measure_layout":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Measure_layout](probe.stageOfInterest)
			for _measure_layout := range set {
				nodeInstance := (&tree.Node{Name: _measure_layout.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_measure_layout, "Measure_layout", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Measure_numbering":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Measure_numbering](probe.stageOfInterest)
			for _measure_numbering := range set {
				nodeInstance := (&tree.Node{Name: _measure_numbering.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_measure_numbering, "Measure_numbering", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Measure_repeat":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Measure_repeat](probe.stageOfInterest)
			for _measure_repeat := range set {
				nodeInstance := (&tree.Node{Name: _measure_repeat.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_measure_repeat, "Measure_repeat", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Measure_style":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Measure_style](probe.stageOfInterest)
			for _measure_style := range set {
				nodeInstance := (&tree.Node{Name: _measure_style.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_measure_style, "Measure_style", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Membrane":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Membrane](probe.stageOfInterest)
			for _membrane := range set {
				nodeInstance := (&tree.Node{Name: _membrane.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_membrane, "Membrane", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Metal":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Metal](probe.stageOfInterest)
			for _metal := range set {
				nodeInstance := (&tree.Node{Name: _metal.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_metal, "Metal", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Metronome":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Metronome](probe.stageOfInterest)
			for _metronome := range set {
				nodeInstance := (&tree.Node{Name: _metronome.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_metronome, "Metronome", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Metronome_beam":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Metronome_beam](probe.stageOfInterest)
			for _metronome_beam := range set {
				nodeInstance := (&tree.Node{Name: _metronome_beam.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_metronome_beam, "Metronome_beam", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Metronome_note":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Metronome_note](probe.stageOfInterest)
			for _metronome_note := range set {
				nodeInstance := (&tree.Node{Name: _metronome_note.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_metronome_note, "Metronome_note", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Metronome_tied":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Metronome_tied](probe.stageOfInterest)
			for _metronome_tied := range set {
				nodeInstance := (&tree.Node{Name: _metronome_tied.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_metronome_tied, "Metronome_tied", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Metronome_tuplet":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Metronome_tuplet](probe.stageOfInterest)
			for _metronome_tuplet := range set {
				nodeInstance := (&tree.Node{Name: _metronome_tuplet.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_metronome_tuplet, "Metronome_tuplet", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Midi_device":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Midi_device](probe.stageOfInterest)
			for _midi_device := range set {
				nodeInstance := (&tree.Node{Name: _midi_device.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_midi_device, "Midi_device", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Midi_instrument":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Midi_instrument](probe.stageOfInterest)
			for _midi_instrument := range set {
				nodeInstance := (&tree.Node{Name: _midi_instrument.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_midi_instrument, "Midi_instrument", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Miscellaneous":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Miscellaneous](probe.stageOfInterest)
			for _miscellaneous := range set {
				nodeInstance := (&tree.Node{Name: _miscellaneous.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_miscellaneous, "Miscellaneous", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Miscellaneous_field":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Miscellaneous_field](probe.stageOfInterest)
			for _miscellaneous_field := range set {
				nodeInstance := (&tree.Node{Name: _miscellaneous_field.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_miscellaneous_field, "Miscellaneous_field", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Mordent":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Mordent](probe.stageOfInterest)
			for _mordent := range set {
				nodeInstance := (&tree.Node{Name: _mordent.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_mordent, "Mordent", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Multiple_rest":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Multiple_rest](probe.stageOfInterest)
			for _multiple_rest := range set {
				nodeInstance := (&tree.Node{Name: _multiple_rest.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_multiple_rest, "Multiple_rest", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Name_display":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Name_display](probe.stageOfInterest)
			for _name_display := range set {
				nodeInstance := (&tree.Node{Name: _name_display.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_name_display, "Name_display", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Non_arpeggiate":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Non_arpeggiate](probe.stageOfInterest)
			for _non_arpeggiate := range set {
				nodeInstance := (&tree.Node{Name: _non_arpeggiate.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_non_arpeggiate, "Non_arpeggiate", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Notations":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Notations](probe.stageOfInterest)
			for _notations := range set {
				nodeInstance := (&tree.Node{Name: _notations.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_notations, "Notations", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Note":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Note](probe.stageOfInterest)
			for _note := range set {
				nodeInstance := (&tree.Node{Name: _note.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_note, "Note", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Note_size":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Note_size](probe.stageOfInterest)
			for _note_size := range set {
				nodeInstance := (&tree.Node{Name: _note_size.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_note_size, "Note_size", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Note_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Note_type](probe.stageOfInterest)
			for _note_type := range set {
				nodeInstance := (&tree.Node{Name: _note_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_note_type, "Note_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Notehead":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Notehead](probe.stageOfInterest)
			for _notehead := range set {
				nodeInstance := (&tree.Node{Name: _notehead.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_notehead, "Notehead", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Notehead_text":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Notehead_text](probe.stageOfInterest)
			for _notehead_text := range set {
				nodeInstance := (&tree.Node{Name: _notehead_text.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_notehead_text, "Notehead_text", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Numeral":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Numeral](probe.stageOfInterest)
			for _numeral := range set {
				nodeInstance := (&tree.Node{Name: _numeral.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_numeral, "Numeral", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Numeral_key":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Numeral_key](probe.stageOfInterest)
			for _numeral_key := range set {
				nodeInstance := (&tree.Node{Name: _numeral_key.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_numeral_key, "Numeral_key", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Numeral_root":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Numeral_root](probe.stageOfInterest)
			for _numeral_root := range set {
				nodeInstance := (&tree.Node{Name: _numeral_root.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_numeral_root, "Numeral_root", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Octave_shift":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Octave_shift](probe.stageOfInterest)
			for _octave_shift := range set {
				nodeInstance := (&tree.Node{Name: _octave_shift.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_octave_shift, "Octave_shift", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Offset":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Offset](probe.stageOfInterest)
			for _offset := range set {
				nodeInstance := (&tree.Node{Name: _offset.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_offset, "Offset", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Opus":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Opus](probe.stageOfInterest)
			for _opus := range set {
				nodeInstance := (&tree.Node{Name: _opus.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_opus, "Opus", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Ornaments":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Ornaments](probe.stageOfInterest)
			for _ornaments := range set {
				nodeInstance := (&tree.Node{Name: _ornaments.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_ornaments, "Ornaments", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Other_appearance":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Other_appearance](probe.stageOfInterest)
			for _other_appearance := range set {
				nodeInstance := (&tree.Node{Name: _other_appearance.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_other_appearance, "Other_appearance", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Other_listening":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Other_listening](probe.stageOfInterest)
			for _other_listening := range set {
				nodeInstance := (&tree.Node{Name: _other_listening.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_other_listening, "Other_listening", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Other_notation":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Other_notation](probe.stageOfInterest)
			for _other_notation := range set {
				nodeInstance := (&tree.Node{Name: _other_notation.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_other_notation, "Other_notation", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Other_play":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Other_play](probe.stageOfInterest)
			for _other_play := range set {
				nodeInstance := (&tree.Node{Name: _other_play.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_other_play, "Other_play", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Page_layout":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Page_layout](probe.stageOfInterest)
			for _page_layout := range set {
				nodeInstance := (&tree.Node{Name: _page_layout.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_page_layout, "Page_layout", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Page_margins":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Page_margins](probe.stageOfInterest)
			for _page_margins := range set {
				nodeInstance := (&tree.Node{Name: _page_margins.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_page_margins, "Page_margins", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Part_clef":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Part_clef](probe.stageOfInterest)
			for _part_clef := range set {
				nodeInstance := (&tree.Node{Name: _part_clef.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_part_clef, "Part_clef", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Part_group":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Part_group](probe.stageOfInterest)
			for _part_group := range set {
				nodeInstance := (&tree.Node{Name: _part_group.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_part_group, "Part_group", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Part_link":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Part_link](probe.stageOfInterest)
			for _part_link := range set {
				nodeInstance := (&tree.Node{Name: _part_link.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_part_link, "Part_link", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Part_list":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Part_list](probe.stageOfInterest)
			for _part_list := range set {
				nodeInstance := (&tree.Node{Name: _part_list.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_part_list, "Part_list", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Part_symbol":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Part_symbol](probe.stageOfInterest)
			for _part_symbol := range set {
				nodeInstance := (&tree.Node{Name: _part_symbol.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_part_symbol, "Part_symbol", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Part_transpose":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Part_transpose](probe.stageOfInterest)
			for _part_transpose := range set {
				nodeInstance := (&tree.Node{Name: _part_transpose.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_part_transpose, "Part_transpose", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Pedal":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Pedal](probe.stageOfInterest)
			for _pedal := range set {
				nodeInstance := (&tree.Node{Name: _pedal.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_pedal, "Pedal", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Pedal_tuning":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Pedal_tuning](probe.stageOfInterest)
			for _pedal_tuning := range set {
				nodeInstance := (&tree.Node{Name: _pedal_tuning.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_pedal_tuning, "Pedal_tuning", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Percussion":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Percussion](probe.stageOfInterest)
			for _percussion := range set {
				nodeInstance := (&tree.Node{Name: _percussion.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_percussion, "Percussion", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Pitch":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Pitch](probe.stageOfInterest)
			for _pitch := range set {
				nodeInstance := (&tree.Node{Name: _pitch.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_pitch, "Pitch", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Pitched":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Pitched](probe.stageOfInterest)
			for _pitched := range set {
				nodeInstance := (&tree.Node{Name: _pitched.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_pitched, "Pitched", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Play":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Play](probe.stageOfInterest)
			for _play := range set {
				nodeInstance := (&tree.Node{Name: _play.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_play, "Play", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Player":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Player](probe.stageOfInterest)
			for _player := range set {
				nodeInstance := (&tree.Node{Name: _player.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_player, "Player", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Principal_voice":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Principal_voice](probe.stageOfInterest)
			for _principal_voice := range set {
				nodeInstance := (&tree.Node{Name: _principal_voice.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_principal_voice, "Principal_voice", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Print":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Print](probe.stageOfInterest)
			for _print := range set {
				nodeInstance := (&tree.Node{Name: _print.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_print, "Print", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Release":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Release](probe.stageOfInterest)
			for _release := range set {
				nodeInstance := (&tree.Node{Name: _release.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_release, "Release", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Repeat":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Repeat](probe.stageOfInterest)
			for _repeat := range set {
				nodeInstance := (&tree.Node{Name: _repeat.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_repeat, "Repeat", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Rest":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Rest](probe.stageOfInterest)
			for _rest := range set {
				nodeInstance := (&tree.Node{Name: _rest.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_rest, "Rest", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Root":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Root](probe.stageOfInterest)
			for _root := range set {
				nodeInstance := (&tree.Node{Name: _root.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_root, "Root", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Root_step":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Root_step](probe.stageOfInterest)
			for _root_step := range set {
				nodeInstance := (&tree.Node{Name: _root_step.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_root_step, "Root_step", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Scaling":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Scaling](probe.stageOfInterest)
			for _scaling := range set {
				nodeInstance := (&tree.Node{Name: _scaling.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_scaling, "Scaling", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Scordatura":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Scordatura](probe.stageOfInterest)
			for _scordatura := range set {
				nodeInstance := (&tree.Node{Name: _scordatura.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_scordatura, "Scordatura", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Score_instrument":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Score_instrument](probe.stageOfInterest)
			for _score_instrument := range set {
				nodeInstance := (&tree.Node{Name: _score_instrument.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_score_instrument, "Score_instrument", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Score_part":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Score_part](probe.stageOfInterest)
			for _score_part := range set {
				nodeInstance := (&tree.Node{Name: _score_part.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_score_part, "Score_part", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Score_partwise":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Score_partwise](probe.stageOfInterest)
			for _score_partwise := range set {
				nodeInstance := (&tree.Node{Name: _score_partwise.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_score_partwise, "Score_partwise", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Score_timewise":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Score_timewise](probe.stageOfInterest)
			for _score_timewise := range set {
				nodeInstance := (&tree.Node{Name: _score_timewise.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_score_timewise, "Score_timewise", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Segno":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Segno](probe.stageOfInterest)
			for _segno := range set {
				nodeInstance := (&tree.Node{Name: _segno.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_segno, "Segno", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Slash":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Slash](probe.stageOfInterest)
			for _slash := range set {
				nodeInstance := (&tree.Node{Name: _slash.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_slash, "Slash", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Slide":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Slide](probe.stageOfInterest)
			for _slide := range set {
				nodeInstance := (&tree.Node{Name: _slide.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_slide, "Slide", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Slur":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Slur](probe.stageOfInterest)
			for _slur := range set {
				nodeInstance := (&tree.Node{Name: _slur.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_slur, "Slur", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Sound":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Sound](probe.stageOfInterest)
			for _sound := range set {
				nodeInstance := (&tree.Node{Name: _sound.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_sound, "Sound", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Staff_details":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Staff_details](probe.stageOfInterest)
			for _staff_details := range set {
				nodeInstance := (&tree.Node{Name: _staff_details.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_staff_details, "Staff_details", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Staff_divide":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Staff_divide](probe.stageOfInterest)
			for _staff_divide := range set {
				nodeInstance := (&tree.Node{Name: _staff_divide.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_staff_divide, "Staff_divide", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Staff_layout":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Staff_layout](probe.stageOfInterest)
			for _staff_layout := range set {
				nodeInstance := (&tree.Node{Name: _staff_layout.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_staff_layout, "Staff_layout", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Staff_size":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Staff_size](probe.stageOfInterest)
			for _staff_size := range set {
				nodeInstance := (&tree.Node{Name: _staff_size.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_staff_size, "Staff_size", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Staff_tuning":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Staff_tuning](probe.stageOfInterest)
			for _staff_tuning := range set {
				nodeInstance := (&tree.Node{Name: _staff_tuning.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_staff_tuning, "Staff_tuning", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Stem":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Stem](probe.stageOfInterest)
			for _stem := range set {
				nodeInstance := (&tree.Node{Name: _stem.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_stem, "Stem", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Stick":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Stick](probe.stageOfInterest)
			for _stick := range set {
				nodeInstance := (&tree.Node{Name: _stick.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_stick, "Stick", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "String_mute":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.String_mute](probe.stageOfInterest)
			for _string_mute := range set {
				nodeInstance := (&tree.Node{Name: _string_mute.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_string_mute, "String_mute", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Strong_accent":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Strong_accent](probe.stageOfInterest)
			for _strong_accent := range set {
				nodeInstance := (&tree.Node{Name: _strong_accent.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_strong_accent, "Strong_accent", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Supports":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Supports](probe.stageOfInterest)
			for _supports := range set {
				nodeInstance := (&tree.Node{Name: _supports.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_supports, "Supports", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Swing":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Swing](probe.stageOfInterest)
			for _swing := range set {
				nodeInstance := (&tree.Node{Name: _swing.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_swing, "Swing", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Sync":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Sync](probe.stageOfInterest)
			for _sync := range set {
				nodeInstance := (&tree.Node{Name: _sync.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_sync, "Sync", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "System_dividers":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.System_dividers](probe.stageOfInterest)
			for _system_dividers := range set {
				nodeInstance := (&tree.Node{Name: _system_dividers.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_system_dividers, "System_dividers", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "System_layout":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.System_layout](probe.stageOfInterest)
			for _system_layout := range set {
				nodeInstance := (&tree.Node{Name: _system_layout.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_system_layout, "System_layout", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "System_margins":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.System_margins](probe.stageOfInterest)
			for _system_margins := range set {
				nodeInstance := (&tree.Node{Name: _system_margins.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_system_margins, "System_margins", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Tap":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Tap](probe.stageOfInterest)
			for _tap := range set {
				nodeInstance := (&tree.Node{Name: _tap.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_tap, "Tap", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Technical":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Technical](probe.stageOfInterest)
			for _technical := range set {
				nodeInstance := (&tree.Node{Name: _technical.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_technical, "Technical", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Text_element_data":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Text_element_data](probe.stageOfInterest)
			for _text_element_data := range set {
				nodeInstance := (&tree.Node{Name: _text_element_data.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_text_element_data, "Text_element_data", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Tie":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Tie](probe.stageOfInterest)
			for _tie := range set {
				nodeInstance := (&tree.Node{Name: _tie.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_tie, "Tie", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Tied":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Tied](probe.stageOfInterest)
			for _tied := range set {
				nodeInstance := (&tree.Node{Name: _tied.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_tied, "Tied", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Time":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Time](probe.stageOfInterest)
			for _time := range set {
				nodeInstance := (&tree.Node{Name: _time.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_time, "Time", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Time_modification":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Time_modification](probe.stageOfInterest)
			for _time_modification := range set {
				nodeInstance := (&tree.Node{Name: _time_modification.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_time_modification, "Time_modification", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Timpani":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Timpani](probe.stageOfInterest)
			for _timpani := range set {
				nodeInstance := (&tree.Node{Name: _timpani.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_timpani, "Timpani", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Transpose":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Transpose](probe.stageOfInterest)
			for _transpose := range set {
				nodeInstance := (&tree.Node{Name: _transpose.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_transpose, "Transpose", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Tremolo":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Tremolo](probe.stageOfInterest)
			for _tremolo := range set {
				nodeInstance := (&tree.Node{Name: _tremolo.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_tremolo, "Tremolo", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Tuplet":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Tuplet](probe.stageOfInterest)
			for _tuplet := range set {
				nodeInstance := (&tree.Node{Name: _tuplet.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_tuplet, "Tuplet", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Tuplet_dot":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Tuplet_dot](probe.stageOfInterest)
			for _tuplet_dot := range set {
				nodeInstance := (&tree.Node{Name: _tuplet_dot.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_tuplet_dot, "Tuplet_dot", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Tuplet_number":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Tuplet_number](probe.stageOfInterest)
			for _tuplet_number := range set {
				nodeInstance := (&tree.Node{Name: _tuplet_number.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_tuplet_number, "Tuplet_number", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Tuplet_portion":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Tuplet_portion](probe.stageOfInterest)
			for _tuplet_portion := range set {
				nodeInstance := (&tree.Node{Name: _tuplet_portion.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_tuplet_portion, "Tuplet_portion", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Tuplet_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Tuplet_type](probe.stageOfInterest)
			for _tuplet_type := range set {
				nodeInstance := (&tree.Node{Name: _tuplet_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_tuplet_type, "Tuplet_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Typed_text":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Typed_text](probe.stageOfInterest)
			for _typed_text := range set {
				nodeInstance := (&tree.Node{Name: _typed_text.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_typed_text, "Typed_text", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Unpitched":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Unpitched](probe.stageOfInterest)
			for _unpitched := range set {
				nodeInstance := (&tree.Node{Name: _unpitched.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_unpitched, "Unpitched", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Virtual_instrument":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Virtual_instrument](probe.stageOfInterest)
			for _virtual_instrument := range set {
				nodeInstance := (&tree.Node{Name: _virtual_instrument.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_virtual_instrument, "Virtual_instrument", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Wait":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Wait](probe.stageOfInterest)
			for _wait := range set {
				nodeInstance := (&tree.Node{Name: _wait.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_wait, "Wait", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Wavy_line":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Wavy_line](probe.stageOfInterest)
			for _wavy_line := range set {
				nodeInstance := (&tree.Node{Name: _wavy_line.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_wavy_line, "Wavy_line", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Wedge":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Wedge](probe.stageOfInterest)
			for _wedge := range set {
				nodeInstance := (&tree.Node{Name: _wedge.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_wedge, "Wedge", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Wood":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Wood](probe.stageOfInterest)
			for _wood := range set {
				nodeInstance := (&tree.Node{Name: _wood.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_wood, "Wood", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Work":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Work](probe.stageOfInterest)
			for _work := range set {
				nodeInstance := (&tree.Node{Name: _work.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_work, "Work", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		}

		nodeGongstruct.IsNodeClickable = true
		nodeGongstruct.Impl = NewTreeNodeImplGongstruct(gongStruct, probe)

		// add add button
		addButton := (&tree.Button{
			Name: gongStruct.Name + " " + string(gongtree_buttons.BUTTON_add),
			Icon: string(gongtree_buttons.BUTTON_add)}).Stage(probe.treeStage)
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)
		addButton.Impl = NewButtonImplGongstruct(
			gongStruct,
			gongtree_buttons.BUTTON_add,
			probe,
		)

		sidebar.RootNodes = append(sidebar.RootNodes, nodeGongstruct)
	}

	// Add a refresh button
	nodeRefreshButton := (&tree.Node{Name: ""}).Stage(probe.treeStage)
	sidebar.RootNodes = append(sidebar.RootNodes, nodeRefreshButton)
	refreshButton := (&tree.Button{
		Name: "RefreshButton" + " " + string(gongtree_buttons.BUTTON_refresh),
		Icon: string(gongtree_buttons.BUTTON_refresh)}).Stage(probe.treeStage)
	nodeRefreshButton.Buttons = append(nodeRefreshButton.Buttons, refreshButton)
	refreshButton.Impl = NewButtonImplRefresh(probe)

	probe.treeStage.Commit()
}

type InstanceNodeCallback[T models.Gongstruct] struct {
	Instance       *T
	gongstructName string
	probe          *Probe
}

func NewInstanceNodeCallback[T models.Gongstruct](
	instance *T,
	gongstructName string,
	probe *Probe) (
	instanceNodeCallback *InstanceNodeCallback[T],
) {
	instanceNodeCallback = new(InstanceNodeCallback[T])

	instanceNodeCallback.probe = probe
	instanceNodeCallback.gongstructName = gongstructName
	instanceNodeCallback.Instance = instance

	return
}

func (instanceNodeCallback *InstanceNodeCallback[T]) OnAfterUpdate(
	gongtreeStage *tree.StageStruct,
	stagedNode, frontNode *tree.Node) {

	FillUpFormFromGongstruct(
		instanceNodeCallback.Instance,
		instanceNodeCallback.probe,
	)
}
