package common

import "time"

var StartTime = time.Now().Unix() // unit: second
var Version = "v1.0.0"            // this hard coding will be replaced automatically when building, no need to manually change

type ModelInfo struct {
	Model     string
	Id        string
	ChuteName string
	MaxTokens int
}

// 创建映射表（假设用 model 名称作为 key）
var ModelRegistry = map[string]ModelInfo{
	"deepseek-r1":                {"deepseek-ai/DeepSeek-R1", "de510462-c319-543b-9c67-00bcf807d2a7", "chutes-deepseek-ai-deepseek-r1", 100000},
	"deepseek-v3-0324":           {"deepseek-ai/DeepSeek-V3-0324", "154ad01c-a431-5744-83c8-651215124360", "chutes-deepseek-ai-deepseek-v3-0324", 100000},
	"deepseek-v3":                {"deepseek-ai/DeepSeek-V3", "fa4181ad-0cf1-5531-afdb-bbffbf2ff945", "chutes-deepseek-ai-deepseek-v3", 100000},
	"qwq-32b":                    {"Qwen/QwQ-32B", "2291db94-1463-5bb3-af2b-72c8d254ee9c", "chutes-qwen-qwq-32", 100000},
	"qwen2.5-72b-instruct":       {"Qwen/Qwen2.5-72B-Instruct", "62cc0462-8983-5ef1-8859-92ccf726e235", "chutes-qwen-qwen2-5-72b-instruct", 100000},
	"qwen2.5-coder-32b-instruct": {"Qwen/Qwen2.5-Coder-32B-Instruct", "2cdc73dd-eec4-5e6b-883a-7736bbe12fd8", "chutes-qwen-qwen2-5-coder-32b-instruct", 100000},
	"gemma-3-27b-it":             {"unsloth/gemma-3-27b-it", "16f41e4f-f2ca-5580-a0a6-46727ae4c212", "chutes-unsloth-gemma-3-27b-it", 100000},
	"olympiccoder-32b":           {"open-r1/OlympicCoder-32B", "649c41b3-25d4-59f9-9124-4d834363ec93", "chutes-open-r1-olympiccoder-32b", 100000},
	"reka-flash-3":               {"RekaAI/reka-flash-3", "893971c6-0f01-54fc-8cbe-420d24a20b54", "chutes-rekaai-reka-flash-3", 100000},
	"ui-tars-72b-dpo":            {"bytedance-research/UI-TARS-72B-DPO", "4d4b0ccf-88a5-5991-a741-df4e99dbf7a2", "chutes-bytedance-research-ui-tars-72b-dpo", 100000},
}

// 通过 model 名称查询的方法
func GetModelInfo(modelName string) (ModelInfo, bool) {
	info, exists := ModelRegistry[modelName]
	return info, exists
}

func GetModelList() []string {
	var modelList []string
	for k := range ModelRegistry {
		modelList = append(modelList, k)
	}
	return modelList
}
