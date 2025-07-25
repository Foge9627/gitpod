/**
 * Copyright (c) 2024 Gitpod GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import { getExperimentsClientForBackend } from "@gitpod/gitpod-protocol/lib/experiments/configcat-server";

export async function getFeatureFlagEnableExperimentalJBTB(userId: string): Promise<boolean> {
    return getExperimentsClientForBackend().getValueAsync("enable_experimental_jbtb", false, {
        user: { id: userId },
    });
}

export async function getFeatureFlagEnableContextEnvVarValidation(userId: string): Promise<boolean> {
    return getExperimentsClientForBackend().getValueAsync("context_env_var_validation", false, {
        user: { id: userId },
    });
}
